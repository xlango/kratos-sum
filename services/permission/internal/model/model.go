package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

type Permission struct {
	Id             int64  `gorm:"primary_key"`
	UserId         string `gorm:"column:UserId;type:int(11);not null;"`
	PermissionName string `gorm:"column:PermissionName;type:varchar(20);not null;"`
}
