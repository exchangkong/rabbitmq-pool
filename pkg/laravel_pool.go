package pool

import (
	"bytes"
	"git.yongche.com/rabbitmq-channel/pb"
	"git.yongche.com/rabbitmq-channel/util"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

type LaravelPool struct {
	PoolService *PoolService
	pb.UnimplementedRabbitmqPoolServiceServer
}

func NewLaravelPoolService(poolService *PoolService) *LaravelPool {
	return &LaravelPool{
		PoolService: poolService,
	}
}

func (l *LaravelPool) Publish(queueName string, content *Content) (msg interface{}, err error) {
	/*
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("laravel publish internal error: %v\n", p)
				return
			}
		}()
	*/

	exchangeName := queueName
	originalQueueName := queueName
	routeKey := queueName

	queueArgs := amqp.Table{}
	if content.Delay != 0 {
		queueArgs["x-dead-letter-exchange"] = ""
		queueArgs["x-dead-letter-routing-key"] = originalQueueName
		queueArgs["x-message-ttl"] = content.Delay * 1000
		queueArgs["x-expires"] = content.Delay * 1000 * 2

		var buffer bytes.Buffer
		buffer.WriteString(queueName)
		buffer.WriteString(".")
		buffer.WriteString("delay")
		buffer.WriteString(".")
		buffer.WriteString(strconv.FormatInt(int64(content.Delay*1000), 10))

		queueName = buffer.String() //延时队列要声明新的延时队列名
		routeKey = queueName
		//只有声明了交换机不是默认交换机的情况下这里才使用交换机
		if content.DeclareExchange {
			queueArgs["x-dead-letter-exchange"] = originalQueueName
		}
	}

	queue := &Queue{
		Name:       queueName,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		Args:       queueArgs,
	}

	exchange := &Exchange{
		Name:     exchangeName,
		Type:     "direct",
		Durable:  true,
		AutoDel:  false,
		Internal: false,
		NoWait:   false,
	}

	retry := 1

	for {
		msg, err = l.PoolService.Publish(queue, exchange, routeKey, content)
		if err == nil {
			break
		}

		retry++
		if retry >= l.PoolService.RetryCount {
			util.FailOnError(err, "pkg php pool retry")
			break
		}
		time.Sleep(time.Millisecond * 30)
	}
	return
}
