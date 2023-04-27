package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zwhypls/grb/api/v1/auth"
)

func main() {
	r := gin.New()
	mc := new(auth.Manager)
	r.POST("/login", mc.SignIn)
	r.Run()
}
