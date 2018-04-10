package main

import (
	pb "go-microservices-tutorials/tutorial2/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"io/ioutil"
	"encoding/json"
	"log"
	//"os"
	"context"
	"github.com/micro/go-micro/cmd"
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
	/*if len(os.Args) > 1 {
		file = os.Args[1]
	}*/

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v",err)
	}

	r,err := client.CreateConsignment(context.TODO(),consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v",err)
	}
	log.Printf("Created: %t",r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v",err)
	}
	for _,v := range getAll.Consignments {
		log.Println(v)
	}
}
