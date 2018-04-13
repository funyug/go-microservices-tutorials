package main

import (
	"os"
	"log"
	"github.com/micro/go-micro"
	vesselProto "github.com/funyug/go-microservices-tutorials/tutorial4/vessel-service/proto/vessel"
	pb "github.com/funyug/go-microservices-tutorials/tutorial4/consignment-service/proto/consignment"
	"fmt"
	"github.com/micro/go-micro/server"
	"context"
	"github.com/micro/go-micro/metadata"
	"errors"
	userService "github.com/funyug/go-microservices-tutorials/tutorial4/user-service/proto/user"
	"github.com/micro/go-micro/client"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session,err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to database with host %s - %v",host,err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run();err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("No auth meta-data found in request")
		}

		token:=meta["Token"]
		log.Println("Authentication with token:", token)

		authClient := userService.NewUserServiceClient("go.micro.srv.user",client.DefaultClient)
		_,err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token:token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req,resp)
		return err
	}
}