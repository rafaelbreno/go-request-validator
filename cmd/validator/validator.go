package validator

import (
	"log"
	"reflect"
)

func Validate(v reflect.Value) {
	typeValid := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeValid.Field(i)
		log.Println("Field:", field.Name, " - Value:", v.Field(i))
		log.Println("-_-_-_-_-_-_-_-_-_-_-_-")
	}
}