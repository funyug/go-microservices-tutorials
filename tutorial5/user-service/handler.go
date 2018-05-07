package main

import (
	"context"
	pb "github.com/funyug/go-microservices-tutorials/tutorial5/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	"log"
	"github.com/micro/go-micro/broker"
	"encoding/json"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
	PubSub		 broker.Broker
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	if err := srv.publishEvent(req); err != nil {
		return err
	}
	return nil
}

func (srv *service) publishEvent(user *pb.User) error {
	body,err := json.Marshal(user)
	if err != nil {
		return err
	}
	msg := &broker.Message{
		Header: map[string]string{
			"id":user.Id,
		},
		Body: body,
	}
	if err := srv.PubSub.Publish(topic,msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	}

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
