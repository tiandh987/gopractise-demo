package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/tiandh987/gopractise-demo/apistyle/greeter/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50023"
	defaultName = "world"
)

func main() {
	// 1. 通过 grpc.Dial(...) 建立一个 gRPC 连接，用来跟服务端进行通信
	// 在创建连接时，可以指定不同的选项，用来控制常见连接的方式。
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	// 2. 连接建立起来之后，需要创建一个客户端 stub，用来执行 RPC 请求
	c := pb.NewGreeterClient(conn)

	// 3. 函数调用前的准备工作
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. 通过 c.SayHello 这种本地式调用方式，调用远端的 SayHello 接口
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
