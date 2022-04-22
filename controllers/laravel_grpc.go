package controllers

import (
	"context"
	"fmt"
	"git.yongche.com/rabbitmq-channel/pb"
	pool "git.yongche.com/rabbitmq-channel/pkg"
	"github.com/streadway/amqp"
)

type GrpcServer struct {
	poolService *pool.LaravelPool
}

func NewGrpcController(pool *pool.LaravelPool) *GrpcServer {
	return &GrpcServer{
		poolService: pool,
	}
}

func(c *GrpcServer) Pushlish(ctx context.Context, request *pb.PublishReuqest) (*pb.PublishResponse, error) {
	response := &pb.PublishResponse{
		Status: 200,
		Message: "",
	}

	fmt.Println("调用poolService.publish")
	publising := &amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(request.Payload),
	}

	content := &pool.Content{
		DeclareExchange: true,
		Delay: request.Delay,
		Publishing: publising,
	}
	_, err := c.poolService.Publish(request.QueueName, content)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
	}

	return response, nil
}