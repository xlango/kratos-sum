package dao

import (
	"context"
	"time"

	kratosRedis "github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/go-redis/redis"
)

//func NewRedis() (r *kratosRedis.Redis, cf func(), err error) {
//	var (
//		cfg kratosRedis.Config
//		ct  paladin.Map
//	)
//	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
//		return
//	}
//	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	r = kratosRedis.NewRedis(&cfg)
//	cf = func() { r.Close() }
//	return
//}

func NewRedis() (r *redis.Client, cf func(), err error) {
	var (
		cfg kratosRedis.Config
		ct  paladin.Map
	)
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		PoolSize:     cfg.Active,
		MinIdleConns: cfg.Idle,
		ReadTimeout:  time.Millisecond * time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(cfg.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
	})
	cf = func() { r.Close() }
	return
}

func (d *dao) SetRedisDB(db int) {
	d.redis.Options().DB = db
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Ping().Result(); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
