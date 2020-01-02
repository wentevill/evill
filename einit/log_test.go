package einit

import (
	"context"
	. "gopkg.in/check.v1"
	"time"

	log "github.com/sirupsen/logrus"
)

func (s *s) TestLog(c *C) {
	initLog()
	log.SetLevel(log.InfoLevel)
	ctx := context.WithValue(context.TODO(), "request_id", "123456a")
	log.WithContext(ctx).Info("this is a info-message")
	log.WithContext(ctx).Warning("this is a waring-message")

	ctx1 := context.WithValue(ctx, "none", "something")
	log.WithContext(ctx1).Info("this is a info-message")

	ctx2, cancel := context.WithDeadline(context.TODO(), time.Now().AddDate(0, 0, 1))
	ctx3 := context.WithValue(ctx2, "request_id", "123456b")
	log.WithContext(ctx3).Info("this is a info-message")
	cancel()
}
