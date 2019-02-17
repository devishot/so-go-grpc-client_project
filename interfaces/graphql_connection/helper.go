package graphql_connection

import (
	"strconv"
	"time"

	conn "github.com/devishot/so-go-grpc-client_project/infrastructure/graphql_connection"
)

func encodeTimestampCursor(t time.Time) conn.Cursor {
	ts := t.Unix()
	str := strconv.FormatInt(ts, 10)
	return conn.NewCursor(str)
}

func decodeTimestampCursor(c conn.Cursor) time.Time {
	str := conn.Must(conn.FromCursor(c)).(string)

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	t := time.Unix(i, 0)
	return t
}
