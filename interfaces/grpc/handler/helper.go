package handler

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

func encodeTimestamp(time time.Time) int64 {
	return time.Unix()
}

func decodeTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func encodeID(id domain.ID) string {
	return string(id)
}

func decodeID(id string) domain.ID {
	return domain.ID(id)
}

func encodeCursor(c domain.Cursor) string {
	return string(c)
}

func encodePageInfo(value conn.PageInfoValue) *pb.PageInfo {
	return &pb.PageInfo{
		HasNextPage:     value.HasNextPage,
		StartCursor:     encodeCursor(value.StartCursor),
		HasPreviousPage: value.HasPreviousPage,
		EndCursor:       encodeCursor(value.EndCursor),
	}
}

func decodeConnRequest(args *pb.ConnArgs) conn.ArgumentsValue {
	switch args.PerPage.(type) {
	case *pb.ConnArgs_First:
		return conn.ArgumentsValue{
			First: int(args.GetFirst()),
			After: domain.Cursor(args.GetAfter()),
		}
	case *pb.ConnArgs_Last:
		return conn.ArgumentsValue{
			Last:   int(args.GetLast()),
			Before: domain.Cursor(args.GetBefore()),
		}
	default:
		panic("proto: ConnArgs.PerPage oneOf-field not defined")
	}
}
