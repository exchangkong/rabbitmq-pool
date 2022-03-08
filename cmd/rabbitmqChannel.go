/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// rabbitmqChannelCmd represents the rabbitmqChannel command
var rabbitmqChannelCmd = &cobra.Command{
	Use:   "rabbitmqChannel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := new(ApplicationOptions)
		opts.Load()

		fmt.Println("rabbitmqChannel called")
		//conn, err := amqp.Dial("amqp://admin:1wbt5NwFEl6cBUoL@172.16.18.8:5672/hashrate")
		conn, err := amqp.Dial("amqp://admin:1wbt5NwFEl6cBUoL@172.16.18.8:5672/hashrate")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		//defer ch.Close()

		/*
		ch1, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch1.Close()
		ch2, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch2.Close()
		*/

		//ch1.Qos(10000, 0, false)
		ch.Confirm(false)

		name := "test_1"
		routeKey := "test_1"
		err = ch.ExchangeDeclare(
			name,   // name
			"direct", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		q, err := ch.QueueDeclare(
			name, // name
			true,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			return
		}
		err = ch.QueueBind(
			q.Name, // queue name
			routeKey,     // routing key
			name, // exchange
			false,
			nil,
		)
		failOnError(err, "Failed to declare a queue")

		body := "Hello World!"
		time.Sleep(3 * time.Second)
		for i := 0; i < 1; i++ {
			err = ch.Publish(
				q.Name,     // exchange
				q.Name, // routing key
				true,   // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			failOnError(err, "faild to queue")
		}

		notifyConfirms := make(chan amqp.Confirmation)
		notifyAck := make(chan uint64)
		notifyNoAck := make(chan uint64)
		/*
		notifyReturn := make(chan amqp.Return)
		notifyCancel := make(chan string)
		*/
		notifyClose := make(chan *amqp.Error)

		ch.NotifyConfirm(notifyAck, notifyNoAck)
		ch.NotifyPublish(notifyConfirms)
		/*
		ch.NotifyReturn(notifyReturn)
		ch.NotifyCancel(notifyCancel)
		*/
		ch.NotifyClose(notifyClose)

		go confirmOnce(notifyAck, notifyNoAck)
		go publishOnce(notifyConfirms)
		/*
		go returnOnce(notifyReturn)
		go cancelOnce(notifyCancel)
		*/
		go closeOnce(notifyClose)

		failOnError(err, "Failed to publish a message")

		/*
		listeners, err := net.Listen("tcp", ":8000")
		if err != nil {
			log.Fatal(err)
		}
		for {
			listeners.Accept()
		}
		*/
		time.Sleep(10*time.Second)
	},
}

func closeOnce(notifyClose <-chan *amqp.Error) {
	for {
		msg := <-notifyClose
		fmt.Println("notifyClose: ", msg)
		return
	}
}

func confirmOnce(notifyAck <-chan uint64, notifyNoAck <-chan uint64) {
	for {
		select {
		case ss := <-notifyAck:
			fmt.Println("confirm notifyAck: ", ss)
			return
		case aa := <-notifyNoAck:
			fmt.Println("confirm notifyNoAck: ", aa)
			return

		}
	}
}

func publishOnce(confirms <-chan amqp.Confirmation) {

	for {
		confirmed := <-confirms
		fmt.Println("publishAck: ", confirmed.Ack)
		fmt.Println("publishDeliveryTag: ", confirmed.DeliveryTag)
		return

		/*
		if confirmed.Ack {
			fmt.Printf("confirmed delivery with delivery tag: %d\n", confirmed.DeliveryTag)
		} else {
			fmt.Printf("publish delivery of delivery tag: %d\n", confirmed.DeliveryTag)
		}
		 */
	}
}

func returnOnce(notifyReturn <-chan amqp.Return) {
	for {
		returns := <-notifyReturn
		fmt.Println("return reply code", returns.ReplyCode)
		fmt.Println("return reply detail", returns)
		return
	}
}

func cancelOnce(notifyCancel <-chan string) {

	for {
		cancel := <-notifyCancel
		fmt.Println("cancel: ", cancel)
		return
	}
}

func init() {
	rootCmd.AddCommand(rabbitmqChannelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rabbitmqChannelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rabbitmqChannelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
