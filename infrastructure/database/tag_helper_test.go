package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractValuesFromTaggedStruct(t *testing.T) {
	type MyORM struct {
		ThisField    string `db:"this_field"`
		AnotherField string `db:"another_field"`
		Timestamp    time.Time
	}

	v := MyORM{"this", "another", time.Now()}
	fields := "this_field, another_field, timestamp"

	values, err := ExtractValuesFromTaggedStruct(fields, reflect.ValueOf(v))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []interface{}{v.ThisField, v.AnotherField, v.Timestamp}, values)
}
