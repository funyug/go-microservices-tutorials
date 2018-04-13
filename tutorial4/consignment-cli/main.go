package main

import (
	pb "github.com/funyug/go-microservices-tutorials/tutorial4/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"context"
	"github.com/micro/go-micro/cmd"
	"errors"
	"github.com/micro/go-micro/metadata"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data,err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data,&consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	file := defaultFilename
	var token string
	log.Println(os.Args)
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v",err)
	}

	ctx := metadata.NewContext(context.Background(),map[string]string{
		"token":token,
	})

	r,err := client.CreateConsignment(ctx,consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v",err)
	}
	log.Printf("Created: %t",r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v",err)
	}
	for _,v := range getAll.Consignments {
		log.Println(v)
	}
}
