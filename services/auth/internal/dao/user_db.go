package dao

import (
	"context"
	"ilinkcloud/services/auth/internal/model"
)

func (d *dao) FindUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {

	db, closeFunc, err := NewDB()
	defer closeFunc()

	resp = new(model.User)
	db = db.Where(&model.User{Username: username}).First(&resp)
	if db.Error != nil {
		err = db.Error
	}
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}
