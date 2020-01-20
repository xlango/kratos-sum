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
	UserId         string `gorm:"type:varchar(11);not null;"`
	PermissionName string `gorm:"type:varchar(20);not null;"`
}
