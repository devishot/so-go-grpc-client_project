package database

import (
	"log"
	"testing"
)

func TestGetValuesInOrder(t *testing.T) {
	type TestStruct struct {
		ID        string
		FirstName string
	}
	expected := []string{"123-456", "Piter Parker"}

	data := TestStruct{expected[0], expected[1]}
	fields := "ID FirstName"

	values, err := GetValuesInOrder(data, fields)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range values {
		expV := expected[i]
		if v != expV {
			log.Fatalf("expected: %v, got: %v", expected, values)
		}
	}
}
