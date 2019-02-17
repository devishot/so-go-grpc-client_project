package handler

import (
	"context"
	"testing"
	"time"

	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

func TestClientService_CreateClient(t *testing.T) {
	s := ClientService{}

	req := &pb.Client{
		Id:          "123",
		CompanyName: "Lamoda",
		Timestamp:   time.Now().Unix(),
	}
	res, err := s.CreateClient(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
