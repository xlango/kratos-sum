package dao

import (
	"context"
	"time"

	pd "Kratos-Grpc-Client/api"
	"Kratos-Grpc-Client/internal/model"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
	SayHello(c context.Context, req *pd.HelloReq) (resp *pd.HelloResp, err error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	demoClient pd.DemoClient
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:         db,
		redis:      r,
		mc:         mc,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	grpcconfig := &warden.ClientConfig{}
	if err = paladin.Get("grpc.toml").UnmarshalTOML(&grpcconfig); err != nil {
		return
	}
	var grpcClient pd.DemoClient
	if grpcClient, err = pd.NewClient(grpcconfig); err == nil {
		d.demoClient = grpcClient
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

func (d *dao) SayHello(c context.Context, req *pd.HelloReq) (resp *pd.HelloResp, err error) {
	c, close := context.WithTimeout(c, 5*time.Second)
	defer close()
	resp, err = d.demoClient.SayHello(c, req)
	return
}
