package controllers

import (
	"context"
	"fmt"
	"git.yongche.com/rabbitmq-channel/pb"
	pool "git.yongche.com/rabbitmq-channel/pkg"
)

type GrpcServer struct {
	poolService *pool.Service
}

func NewGrpcController(pool *pool.Service) *GrpcServer {
	return &GrpcServer{
		poolService: pool,
	}
}

func(c *GrpcServer) Pushlish(ctx context.Context, request *pb.PublishReuqest) (*pb.PublishResponse, error) {
	content := &pool.Content{
		QueueName: request.QueueName,
		DeclareExchange: request.DeclareExchange,
		ContentType: "application/json",
		Body: request.Payload,
		Delay: request.Delay,
	}

	fmt.Println("调用poolService.publish")
	c.poolService.Publish(content)
}