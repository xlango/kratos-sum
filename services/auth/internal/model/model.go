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

type User struct {
	Username string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"type:varchar(20);not null;"`
}
