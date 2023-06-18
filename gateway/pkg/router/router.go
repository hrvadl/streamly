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
	r.POST("/sign-up", authSVC.HandleSignUp())

	r.GET("/user/:id", userSVC.HandleGetByID())
	r.POST("/user", userSVC.HandleCreate())

	return r
}
