package main

import (
	pb "github.com/funyug/go-microservices-tutorials/tutorial3/vessel-service/proto/vessel"
	"errors"
	"context"
	"github.com/micro/go-micro"
	"fmt"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel,error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel,error){
	for _,vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel,nil
		}
	}
	return nil, errors.New("No vessel found by spec")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel,err:=s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id:"vessel01",Name:"Boaty McBoatface",MaxWeight:2000,Capacity:500},
	}
	repo:=&VesselRepository{vessels}

	srv:=micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(),&service{repo})

	if err:=srv.Run();err != nil {
		fmt.Println(err)
	}
}