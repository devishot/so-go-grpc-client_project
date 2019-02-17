package domain

import (
	"encoding/base64"
	"strconv"
	"time"
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

func EncodeTimestampCursor(t time.Time) Cursor {
	ts := t.Unix()
	str := strconv.FormatInt(ts, 10)
	return NewCursor(str)
}

func DecodeTimestampCursor(c Cursor) time.Time {
	str := Must(FromCursor(c)).(string)

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	t := time.Unix(i, 0)
	return t
}

func Must(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}
