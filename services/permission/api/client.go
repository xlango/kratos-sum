package api

import (
	"context"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "direct://default/127.0.0.1:9001"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), AppID)
	if err != nil {
		return nil, err
	}
	return NewDemoClient(cc), nil
}

// target server addrs.

//const target = "direct://default/127.0.0.1:9001" // NOTE: example
//
//// NewClient new member grpc client
//func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
//	client := warden.NewClient(cfg, opts...)
//	conn, err := client.Dial(context.Background(), target)
//	if err != nil {
//		return nil, err
//	}
//	// 注意替换这里：
//	// NewDemoClient方法是在"api"目录下代码生成的
//	// 对应proto文件内自定义的service名字，请使用正确方法名替换
//	return NewDemoClient(conn), nil
//}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
