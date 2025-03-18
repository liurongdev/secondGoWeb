package client

import (
	"context"
	"fmt"
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
	//defer func() {
	//	r := recover()
	//	if r != nil {
	//		fmt.Println(r)
	//	}
	//}()

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	// 将 metadata 添加到上下文中
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 调用远程方法
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		fmt.Printf("SayHello err:%v\n", err)
	}

	log.Printf("Response from server: %s", resp)
	return resp.Message
}
