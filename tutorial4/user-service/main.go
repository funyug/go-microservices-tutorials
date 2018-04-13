package main

import (
	pb "github.com/funyug/go-microservices-tutorials/tutorial4/user-service/proto/user"
	"log"
	"github.com/micro/go-micro"
	"fmt"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo,tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
