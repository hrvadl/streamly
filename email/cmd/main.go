package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getbrevo/brevo-go/lib"
	"github.com/hrvadl/studdy-buddy/email/pkg/config"
	"github.com/hrvadl/studdy-buddy/email/pkg/service"
)

func main() {
	l := log.Default()
	cfg := config.Load()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		l.Fatalf("cannot connect to kafka %v", err)
	}

	l.Println("created email consumer on URL ")

	email := service.InitEmail(cfg)
	email.SendEmail(&lib.SendSmtpEmail{})
}
