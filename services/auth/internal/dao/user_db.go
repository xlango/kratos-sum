package dao

import (
	"context"
	"ilinkcloud/services/auth/internal/model"
)

func (d *dao) FindUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {

	resp = new(model.User)
	d.db = d.db.Where(&model.User{Username: username}).First(&resp)
	if d.db.Error != nil {
		err = d.db.Error
	}
	return
}

func (d *dao) UpdateUserByUsername(ctx context.Context, username string, req *model.User) (err error) {

	err = d.db.Table("tb_user").Where("username = ?", username).UpdateColumn("password", req.Password).Error
	if d.db.Error != nil {
		err = d.db.Error
	}
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}
