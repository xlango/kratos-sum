package dao

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/jinzhu/gorm"
	"ilinkcloud/services/auth/internal/model"
)

//func NewDB() (db *sql.DB, cf func(), err error) {
//	var (
//		cfg sql.Config
//		ct  paladin.TOML
//	)
//	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
//		return
//	}
//	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	db = sql.NewMySQL(&cfg)
//	cf = func() { db.Close() }
//	return
//}

type MysqlCfg struct {
	sql.Config
	Prefix string
}

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg MysqlCfg
		ct  paladin.TOML
	)

	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}

	//连接串
	db, err = gorm.Open("mysql", cfg.DSN)

	//defer db.Close()
	if err != nil {
		panic(err)
		return
	}
	cf = func() { db.Close() }
	//设置最大空闲连接数和最大连接
	db.DB().SetMaxIdleConns(cfg.Idle)
	db.DB().SetMaxOpenConns(cfg.Active)
	//true:不使用结构体名称的复数形式映射生成表名
	db.SingularTable(true)
	//设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return cfg.Prefix + defaultTableName
	}

	return
}

func InitTable() {
	CreateTalbe(model.User{})
}

func CreateTalbe(v interface{}) {
	msdb, closeFunc, err := NewDB()
	if err != nil {
		panic(err)
	}
	defer closeFunc()
	//判断表是否存在，不存在则创建
	if !msdb.HasTable(v) {
		if err := msdb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(v).Error; err != nil {
			panic(err)
		}
	}
}
