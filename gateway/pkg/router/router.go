package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/auth"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/user"
)

func Configure(c *config.Config) *gin.Engine {
	r := gin.Default()

	authSVC := auth.NewService(auth.InitClient(c))
	userSVC := user.NewService(user.InitClient(c))

	r.POST("/sign-in", authSVC.HandleSignIn())
	r.GET("/user/:id", userSVC.HandleGetByID())

	return r
}
