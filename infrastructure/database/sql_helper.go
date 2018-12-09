package database

import (
	"io"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustClose(rows io.Closer) {
	Must(rows.Close())
}
