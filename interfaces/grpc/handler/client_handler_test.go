package handler

import (
	"context"
	"testing"
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
	conn_services "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/services"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

func getTestClientService() ClientService {
	clConnRepo := &repository.ClientConnectionRepositoryMock{}
	clRepo := &domain.ClientRepositoryMock{}

	return ClientService{
		ConnBuilder: conn_services.Builder{ClientConnRepo: clConnRepo},
		Service:     domain.ClientService{Repo: clRepo},
	}
}

func TestClientService_CreateClient(t *testing.T) {
	var resClient domain.ClientEntity

	s := getTestClientService()
	s.Service.Repo = &domain.ClientRepositoryMock{
		CreateFunc: func(entity domain.ClientEntity) error {
			resClient = entity
			return nil
		},
	}

	data := &pb.Client{
		CompanyName: "Lamoda",
	}
	req := &pb.CreateClientRequest{
		Data: data,
	}
	resp, err := s.CreateClient(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resClient.IsZero() {
		t.Fatalf("expected: %v, got: %v", data, resClient)
	}
	if resp.Id != string(resClient.ID) ||
		resp.CompanyName != resClient.CompanyName ||
		resp.Timestamp == 0 {
		t.Fatalf("grpc data not equal to domain data: %v %v", resp, data)
	}

	t.Log(resp)
}

func TestClientService_GetClientConnection(t *testing.T) {
	last := domain.NewClient("", "", "Trump Empire")
	lastCursor := domain.EncodeTimestampCursor(last.Timestamp)
	cursor := domain.EncodeTimestampCursor(time.Now())
	perPage := 10

	s := getTestClientService()
	s.ConnBuilder.ClientConnRepo = &repository.ClientConnectionRepositoryMock{
		GetLastFunc: func() (entity domain.ClientEntity, e error) {
			return last, nil
		},
		PaginateForwardByTimestampFunc: func(first int, after time.Time) (entities []domain.ClientEntity, e error) {
			if first != perPage {
				t.Fatalf("expected: %v, got: %v", perPage, first)
			}

			return []domain.ClientEntity{last}, nil
		},
	}

	req := &pb.GetClientConnectionRequest{
		Args: &pb.ConnArgs{
			PerPage:        &pb.ConnArgs_First{First: int32(perPage)},
			PreviousCursor: &pb.ConnArgs_After{After: string(cursor)},
		},
	}
	resp, err := s.GetClientConnection(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.PageInfo.HasNextPage != false {
		t.Fatalf("expected: no any edges after last page")
	}
	if resp.PageInfo.EndCursor != string(lastCursor) {
		t.Fatalf("expected: %s, got: %s", lastCursor, resp.PageInfo.EndCursor)
	}
	if resp.Edges[0].Cursor != string(lastCursor) {
		t.Fatalf("expected: %s, got: %s", lastCursor, resp.Edges[0].Cursor)
	}

	t.Log(resp)
}
