package main

import (
	pb "github.com/funyug/go-microservices-tutorials/tutorial5/user-service/proto/user"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"log"
	"github.com/micro/go-micro/broker"
	"encoding/json"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error {
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
