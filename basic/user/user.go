package main

import (
	"context"
	"evill/basic/user/model"
	"evill/basic/user/proto"

	"github.com/go-xorm/xorm"
)

type User struct {
	db *xorm.EngineGroup
}

func (u *User) SignUp(ctx context.Context, in *user.SignUpRequest) (out *user.SignUpResponse, err error) {
	out = new(user.SignUpResponse)
	user := &model.User{
		UserName: in.Name,
	}
	if _, err := u.db.Table(user).Insert(user); err != nil {
		return out, err
	}
	return
}

func (u *User) SignIn() {
}

func (u *User) Cancellation() {

}

func (u *User) ModifyBasic() {

}
