package main

import (
	"context"
	user "evill/basic/user/proto"
	"evill/einit"
	. "gopkg.in/check.v1"
	"testing"
)

func TestAll(t *testing.T) {
	TestingT(t)
}

type s struct {
}

var _ = Suite(&s{})

func (s *s) TestInit(c *C) {
	res, err := einit.Init(einit.Log|einit.Mysql, "../../einit/config.yml")
	c.Check(err, IsNil)
	c.Check(res.Mysql(), NotNil)
	u := &User{
		db: res.Mysql(),
	}
	ctx := context.TODO()
	resp, err := u.SignUp(ctx, &user.SignUpRequest{
		Name: "insert two",
	})
	c.Check(err, IsNil)
	c.Check(resp, IsNil)
}
