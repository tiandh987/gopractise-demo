package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/tiandh987/gopractise-demo/apistyle/greeter/helloworld"
)

// 1. 定义一个 server 结构体
type server struct {
	pb.UnimplementGreeterServer
}

// 2. 为 server 结构体添加 SayHello 方法，
// 也就是说 server 是 GreeterServer 接口的一个实现。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

const (
	port = ":50023"
)

func main() {
	// 3. 通过 net.Listen(...) 指定监听客户端请求的端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 4. 通过 grpc.NewServer() 创建一个 gRPC Server 实例
	s := grpc.NewServer()

	// 5. 将该服务注册到 gRPC 框架中
	pb.RegisterGreeterServer(s, &server{})

	// 6. 启动 gRPC 服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
