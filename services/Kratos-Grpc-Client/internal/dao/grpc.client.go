package dao

import (
	pb "Kratos-Grpc-Client/api"
	"context"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

//定义target 为gRPC用于服务发现的目标
const target = "direct://default/127.0.0.1:11000"

func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (pb.DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), target)
	if err != nil {
		return nil, err
	}

	return pb.NewDemoClient(conn), nil
}
