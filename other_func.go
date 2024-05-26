package main

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strconv"
)

func StructToMapStringString(data interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()

		switch v := fieldValue.(type) {
		case string:
			result[fieldName] = v
		case int:
			result[fieldName] = strconv.Itoa(v)
		case bool:
			result[fieldName] = strconv.FormatBool(v)
		case float64:
			result[fieldName] = strconv.FormatFloat(v, 'f', -1, 64)
		default:
			result[fieldName] = fmt.Sprintf("%v", v)
		}
	}

	return result
}

func LogError(err error, message string) {
	if err != nil {
		slog.Error(message)
	}
}

func CheckAndCreateLogJSON() {
	_, err := os.Stat("log.json")
	if err != nil {
		f, err := os.Create("log.json")
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}
