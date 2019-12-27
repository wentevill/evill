package model

type User struct {
	Id       int64
	UserName string
	Avatar   string
	Email    string
	CreateAt int64 `xorm:"created"`
	UpdateAt int64 `xorm:"updated"`
	DeleteAt int64 `xorm:"deleted"`
}
