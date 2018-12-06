package database

import (
	"strings"

	// TODO: replace this library,
	//  it has incorrect import for tests.
	"gopkg.in/oleiade/reflections.v1"
)

func GetValuesInOrder(data interface{}, fields string) (result []interface{}, err error) {
	fieldsArr := strings.Split(fields, " ")

	for _, fieldName := range fieldsArr {
		fieldValue, fieldErr := reflections.GetField(data, fieldName)
		if fieldErr != nil {
			err = fieldErr
			return
		}

		result = append(result, fieldValue)
	}

	return
}
