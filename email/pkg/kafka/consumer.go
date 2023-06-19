package kafka

import (
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getbrevo/brevo-go/lib"
	"github.com/hrvadl/studdy-buddy/email/pkg/config"
	"github.com/hrvadl/studdy-buddy/email/pkg/pb"
	"github.com/hrvadl/studdy-buddy/email/pkg/service"
	"google.golang.org/protobuf/proto"
)

const checkDelay = time.Second

func InitConsumer(cfg *config.Config, l *log.Logger) {
	sigchan := make(chan os.Signal, 1)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.BootstrapServers,
		"group.id":                 cfg.GroupID,
		"auto.offset.reset":        "earliest",
		"session.timeout.ms":       6000,
		"allow.auto.create.topics": true,
	})

	if err != nil {
		l.Fatalf("cannot connect to kafka %v", err)
	}

	consumer.SubscribeTopics([]string{"sendEmail"}, nil)
	defer consumer.Close()
	l.Println("created email consumer")

	email := service.InitEmail(cfg)

cons:
	for {
		select {
		case sig := <-sigchan:
			l.Printf("received signal: %v\n", sig)
			break cons

		default:
			evt, err := consumer.ReadMessage(checkDelay)

			if err != nil {
				l.Print(err)
				continue
			}

			if evt == nil {
				continue
			}

			var mp pb.MailPayload

			if err := proto.Unmarshal(evt.Value, &mp); err != nil {
				l.Printf("error unmarshalling message: %v", err)
				continue
			}

			_, err = email.SendEmail(&lib.SendSmtpEmail{})

			if err != nil {
				l.Printf("failed to send email: %v\n", err)
			}
		}
	}
}
