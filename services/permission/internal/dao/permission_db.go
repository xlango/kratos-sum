package dao

import (
	"context"
	"ilinkcloud/services/permission/internal/model"
)

func (d *dao) AddPermission(ctx context.Context, req *model.Permission) (err error) {

	err = d.db.Create(&req).Error
	if d.db.Error != nil {
		err = d.db.Error
	}
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}
