package client

import (
	"awesomeProject2/middleware/logger"
	"context"
	pb "github.com/liurongdev/firstGoWeb/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"time"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
}

func TestCallRemoteGrpc() string {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			logger.Info("error_result:", err)
		}
	}()
	client := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用远程方法
	logger.Info("client.sayHello")
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		logger.Error("SayHello err:%v\n", err)
	}
	log.Printf("Response from server: %s", resp)
	return resp.Message
}
