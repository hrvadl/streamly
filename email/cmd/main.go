package main

import (
	"log"

	"github.com/hrvadl/studdy-buddy/email/pkg/config"
	"github.com/hrvadl/studdy-buddy/email/pkg/kafka"
)

func main() {
	l := log.Default()
	cfg := config.Load()
	kafka.InitConsumer(cfg, l)
}
