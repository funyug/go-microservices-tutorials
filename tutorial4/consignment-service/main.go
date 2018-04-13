package main

import (
	"os"
	"log"
	"github.com/micro/go-micro"
	vesselProto "github.com/funyug/go-microservices-tutorials/tutorial4/vessel-service/proto/vessel"
	pb "github.com/funyug/go-microservices-tutorials/tutorial4/consignment-service/proto/consignment"
	"fmt"
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
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run();err != nil {
		fmt.Println(err)
	}
}
