package handler

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

func encodeProtoClient(cl domain.ClientEntity) *pb.Client {
	return &pb.Client{
		Id:          string(cl.ID),
		Timestamp:   encodeTimestamp(cl.Timestamp),
		FirstName:   cl.FirstName,
		LastName:    cl.LastName,
		CompanyName: cl.CompanyName,
	}
}

func decodeProtoClient(cl *pb.Client) domain.ClientEntity {
	return domain.ClientEntity{
		ID:          decodeID(cl.Id),
		Timestamp:   decodeTimestamp(cl.Timestamp),
		FirstName:   cl.FirstName,
		LastName:    cl.LastName,
		CompanyName: cl.CompanyName,
	}
}

func encodeClientConnResponse(value conn.PageValue) *pb.ClientConnectionResponse {
	edges := make([]*pb.ClientConnectionEdge, 0)
	for _, val := range value.Edges {
		cl, _ := val.Node.(domain.ClientEntity)

		edges = append(edges, &pb.ClientConnectionEdge{
			Cursor: encodeCursor(val.Cursor),
			Node:   encodeProtoClient(cl),
		})
	}

	return &pb.ClientConnectionResponse{
		Edges:    edges,
		PageInfo: encodePageInfo(value.PageInfo),
	}
}
