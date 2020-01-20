package dao

import (
	"context"
	"github.com/prometheus/common/log"
	pb "ilinkcloud/services/permission/api"
	"ilinkcloud/services/permission/internal/dao/tran"
	"ilinkcloud/services/permission/internal/model"
)

type PermissionDao interface {
	PermissionSave(ctx context.Context, req *pb.PermissionSaveReq) (resp *pb.PermissionSaveResp, err error)
}

func (d *dao) PermissionSave(ctx context.Context, req *pb.PermissionSaveReq) (resp *pb.PermissionSaveResp, err error) {
	groupId := ctx.Value("tranGroupId")

	tx, err := tran.TMBegin(d.db, false)
	tx.Msg.GroupId = groupId.(string)
	defer tx.Close()
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	p := model.Permission{
		UserId:         req.UserId,
		PermissionName: req.PermissionName,
	}

	err = d.AddPermission(ctx, &p)
	if err != nil {
		log.Errorln(err)
		tx.RMRollback(true)
	} else {
		tx.RMCommit(true)
	}

	tx.Commit()

	resp = new(pb.PermissionSaveResp)
	resp.UserId = req.UserId
	//resp.PermissionId=req.Password

	return
}
