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
	"git.yongche.com/rabbitmq-channel/pkg"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := new(ApplicationOptions)
		opts.Load()

		viper.OnConfigChange(func(e fsnotify.Event) {
			// 配置文件发生变更之后会调用的回调函数
			fmt.Println("Config file changed:", e.Name)
			opts.Load()
			fmt.Println(opts.PoolOption)
		})
		boot := newBootStrap(opts)
		poolService := boot.PoolService

		laravelPool := pool.NewLaravelPoolService(poolService)

		fmt.Println("pool called")

		http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
			values := request.URL.Query()
			data := values.Get("data")
			delay, _ := strconv.ParseInt(values.Get("delay"), 10, 32)

			content := &pool.Content{
				DeclareExchange: true,
				ContentType: "text/plain",
				Body: data,
				Delay: int32(delay),
			}

			numbers, _ := strconv.ParseInt(values.Get("num"), 10, 32)

			for i := int64(1); i <= numbers; i++ {
				_, _ =laravelPool.Publish("hello", content)
			}

			writer.WriteHeader(200)
			writer.Write([]byte("阿斯顿发斯蒂芬"))
		})
		http.ListenAndServe(":8080", nil)

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)

		<-quit
	},
}

func init() {
	rootCmd.AddCommand(poolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

