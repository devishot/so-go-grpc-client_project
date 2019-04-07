package handler

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

func encodeProtoClientProject(cl domain.ProjectEntity) *pb.ClientProject {
	return &pb.ClientProject{
		Id:          encodeID(cl.ID),
		Timestamp:   encodeTimestamp(cl.Timestamp),
		Title:       cl.Title,
		Description: cl.Description,
		ClientId:    encodeID(cl.ClientID),
	}
}

func decodeProtoClientProject(cl *pb.ClientProject) domain.ProjectEntity {
	return domain.ProjectEntity{
		ID:          decodeID(cl.Id),
		Timestamp:   decodeTimestamp(cl.Timestamp),
		Title:       cl.Title,
		Description: cl.Description,
		ClientID:    decodeID(cl.ClientId),
	}
}

func encodeClientProjectConnResponse(value conn.PageValue) *pb.ClientProjectConnectionResponse {
	edges := make([]*pb.ClientProjectConnectionEdge, 0)
	for _, val := range value.Edges {
		p, _ := val.Node.(domain.ProjectEntity)

		edges = append(edges, &pb.ClientProjectConnectionEdge{
			Cursor: encodeCursor(val.Cursor),
			Node:   encodeProtoClientProject(p),
		})
	}

	return &pb.ClientProjectConnectionResponse{
		Edges:    edges,
		PageInfo: encodePageInfo(value.PageInfo),
	}
}
