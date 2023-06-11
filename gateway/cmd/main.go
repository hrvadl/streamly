package main

import (
	"fmt"

	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/router"
)

func main() {
	c := config.Load()
	r := router.Configure(c)
	r.Run(fmt.Sprintf(":%v", c.Port))
}
