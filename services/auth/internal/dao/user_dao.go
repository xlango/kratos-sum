package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	pb "ilinkcloud/services/auth/api"
	"ilinkcloud/services/auth/internal/model"
	permissionpb "ilinkcloud/services/permission/api"
	"ilinkcloud/services/public/tran"
	"math/rand"
	"time"
)

type UserDao interface {
	Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error)
	UserSave(ctx context.Context, req *pb.UserSaveReq) (resp *pb.UserSaveResp, err error)
}

func (d *dao) Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error) {
	user, err := d.FindUserByUsername(ctx, req.Username)
	defer d.Close()
	if err != nil {
		err = errors.New("username is not exist")
		return
	}

	if user.Password != req.Password {
		err = errors.New("password error")
		return
	}

	resp = new(pb.UserLoginResp)
	resp.Username = user.Username
	resp.Token = GetRandomString(15)

	//err = d.SaveUserToken(resp.Token, resp.Username)
	//if err != nil {
	//	err = errors.New("token create failed")
	//	return
	//}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	_, err = d.permissionClient.SayHello(ctx, &permissionpb.HelloReq{Name: resp.Token})

	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (d *dao) UserSave(ctx context.Context, req *pb.UserSaveReq) (resp *pb.UserSaveResp, err error) {
	tx, err := tran.TMBegin(d.db, true, 2)
	defer tx.Close()

	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	err = d.UpdateUserByUsername(ctx, req.Username, &model.User{Password: req.Password})

	ctx = context.WithValue(ctx, "tranGroupId", tx.Msg.GroupId)

	if err != nil {
		log.Errorln(err)
		tx.RMRollback(true)
	} else {
		tx.RMCommit(true)
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	_, err = d.permissionClient.PermissionSave(ctx, &permissionpb.PermissionSaveReq{UserId: req.Username, PermissionName: "per1"})
	if err != nil {
		log.Errorln(err)
		tx.TMCancel()
	}

	tx.Commit()
	resp = new(pb.UserSaveResp)
	resp.Username = req.Username
	resp.Password = req.Password
	return
}
