package einit

import (
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
	res, err := Init(Mysql, "./config.yml")
	c.Check(err, IsNil)
	c.Check(res[Mysql], NotNil)
}
