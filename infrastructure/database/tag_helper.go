package database

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
)

const TagName = "db"
const FieldsSeparator = ", "

func ExtractValuesFromTaggedStruct(obj interface{}, fields string) (arr []interface{}, err error) {
	v := reflect.ValueOf(obj)
	tags := ParseFields(fields)

	m := reflectx.NewMapperTagFunc(TagName, strings.ToLower, nil)

	for _, tag := range tags {
		valueByTag := m.FieldByName(v, tag)

		if !valueByTag.IsValid() || valueByTag.Type() == v.Type() {
			err = errors.Errorf("cannot find tag=%s in struct (type=%v, value=%v)", tag, v.Type(), v.Interface())
			err = errors.WithStack(err)
			return
		}

		val := valueByTag.Interface()
		arr = append(arr, val)
	}

	return
}

func ParseFields(fields string) []string {
	return strings.Split(fields, FieldsSeparator)
}
