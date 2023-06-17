package main

import (
	"fmt"
	"log"

	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/router"
)

func main() {
	l := log.Default()
	l.Println("starting gateway...")
	c := config.Load()
	r := router.Configure(c)
	l.Printf("running gateway on port %v", c.Port)

	if err := r.Run(fmt.Sprintf(":%v", c.Port)); err != nil {
		l.Fatalf("cannot start gateway: %v", err)
	}
}
