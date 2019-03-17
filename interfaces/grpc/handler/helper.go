package handler

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
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

func encodeTimestamp(time time.Time) int64 {
	return time.Unix()
}

func decodeTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func decodeID(id string) domain.ID {
	return domain.ID(id)
}

func encodeCursor(c domain.Cursor) string {
	return string(c)
}

func encodeClientConnResponse(value domain.ConnectionValue) *pb.ClientConnectionResponse {
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

func encodePageInfo(value domain.ConnectionPageInfoValue) *pb.PageInfo {
	return &pb.PageInfo{
		HasNextPage:     value.HasNextPage,
		StartCursor:     encodeCursor(value.StartCursor),
		HasPreviousPage: value.HasPreviousPage,
		EndCursor:       encodeCursor(value.EndCursor),
	}
}

func decodeConnRequest(args *pb.ConnArgs) domain.ConnectionArgumentsValue {
	switch args.PerPage.(type) {
	case *pb.ConnArgs_First:
		return domain.ConnectionArgumentsValue{
			First: int(args.GetFirst()),
			After: domain.Cursor(args.GetAfter()),
		}
	case *pb.ConnArgs_Last:
		return domain.ConnectionArgumentsValue{
			Last:   int(args.GetLast()),
			Before: domain.Cursor(args.GetBefore()),
		}
	default:
		panic("proto: ConnArgs.PerPage oneOf-field not defined")
	}
}
