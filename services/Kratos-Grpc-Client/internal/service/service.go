package service

import (
	"context"
	"fmt"

	pb "Kratos-Grpc-Client/api"
	"Kratos-Grpc-Client/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	span1, ctx := opentracing.StartSpanFromContext(ctx, "Client-SayHello")
	defer span1.Finish()
	reply, err = s.dao.SayHello(ctx, req)
	if err != nil {
		reply = &pb.HelloResp{
			Content: err.Error(),
		}
	}
	//reply = &pb.HelloResp{}
	reply.Content = reply.Content + "-----Client"
	return
}

// SayHelloURL bm demo func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	fmt.Printf("hello url %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
