package dao

import (
	"context"
	"errors"
	"fmt"
	pb "ilinkcloud/services/auth/api"
	permissionpb "ilinkcloud/services/permission/api"
	"math/rand"
	"time"
)

type UserDao interface {
	Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error)
}

func (d *dao) Login(ctx context.Context, req *pb.UserLoginReq) (resp *pb.UserLoginResp, err error) {
	user, err := d.FindUserByUsername(ctx, req.Username)
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

	//ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	//defer cancel()

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
