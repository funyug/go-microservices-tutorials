package main

import (
	pb "github.com/funyug/go-microservices-tutorials/tutorial4/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"context"
	"log"
	"os"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user",microclient.DefaultClient)
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name: "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name: "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:"password",
				Usage:"Your password",
			},
			cli.StringFlag{
				Name:"company",
				Usage:"Your company",
			},
		),
	)

	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			r, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v",err)
			}
			log.Printf("Created: %s",r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v",err)
			}
			for _,v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
