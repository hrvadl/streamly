package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
)

func main() {
	r := gin.Default()
	c := config.Load()
	r.Run(":5000")
}
