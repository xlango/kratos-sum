package service

import (
	"context"
	pb "ilinkcloud/services/auth/api"
)

func (s *Service) Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error) {
	resp, err = s.dao.Login(ctx, req)
	if err != nil {
		resp = new(pb.UserLoginResp)
		resp.Token = err.Error()
	}
	return
}

func (s *Service) UserSave(ctx context.Context, req *pb.UserSaveReq) (resp *pb.UserSaveResp, err error) {
	resp, err = s.dao.UserSave(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}
