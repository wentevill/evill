package main

import (
	"evill/einit"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func newRouter() {
	r := gin.Default()

	v1 := r.Group("/v1")

	register(new(user), v1)

	_ = r.Run(einit.GetConfig().Port)
}

func register(b base, v ...*gin.RouterGroup) {
	if err := b.init(); err != nil {
		log.Error(err)
		return
	}
	b.register(v...)
}

type base interface {
	register(...*gin.RouterGroup)
	name() string
	init() error
}

func errResponse(c *gin.Context, err error) {
	log.Error(err)
	resp(c, http.StatusInternalServerError, nil)
}
func response(c *gin.Context, data interface{}) {
	resp(c, http.StatusOK, data)
}

func resp(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
}
