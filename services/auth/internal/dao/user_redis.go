package dao

import "fmt"

type UserRedis interface {
	GetUserRedis() error
}

func keyRedis(key string) string {
	return fmt.Sprintf("User:Token:%v", key)
}

func (d *dao) SaveUserToken(key string, val interface{}) (err error) {
	d.SetRedisDB(0)
	err = d.redis.Set(keyRedis(key), val, 0).Err()

	return
}
