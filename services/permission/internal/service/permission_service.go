package service

import (
	"context"
	pb "ilinkcloud/services/permission/api"
)

func (s *Service) PermissionSave(ctx context.Context, req *pb.PermissionSaveReq) (resp *pb.PermissionSaveResp, err error) {
	resp = new(pb.PermissionSaveResp)
	resp, err = s.dao.PermissionSave(ctx, req)
	return
}
