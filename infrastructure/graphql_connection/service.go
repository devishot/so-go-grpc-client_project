package graphql_connection

import (
	"encoding/base64"
)

func NewCursor(str string) Cursor {
	cursor := base64.StdEncoding.EncodeToString([]byte(str))
	return Cursor(cursor)
}

func FromCursor(c Cursor) (string, error) {
	b, err := base64.StdEncoding.DecodeString(string(c))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Must(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}
