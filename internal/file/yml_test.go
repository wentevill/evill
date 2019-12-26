package file

import (
	"gopkg.in/check.v1"
	"testing"
)

func TestAll(t *testing.T) {
	check.TestingT(t)
}

type s struct {
}

var _ = check.Suite(&s{})

func (s *s) TestLoad(c *check.C) {
	path := "test.yml"
	var t = struct {
		Level int `yml:"level"`
	}{}
	err := new(YML).Load(&t, path)
	c.Check(err, check.IsNil)
	c.Check(t.Level, check.Equals, 1)
}
