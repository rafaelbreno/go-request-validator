package validator

import (
	"log"
	"strings"
	"reflect"
)

func Validate(v reflect.Value) {
	typeValid := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeValid.Field(i)
		// log.Println(field.Tag)
		
		for ruleKey, ruleValue := range filterRules(string(field.Tag)) {
			log.Println(ruleKey, ruleValue)
		}
		// log.Println("Field:", field.Name, " - Value:", v.Field(i))
		// log.Println("-_-_-_-_-_-_-_-_-_-_-_-")
	}
}

func filterRules(rawRules string) map[string]string {
	rules := strings.Split(strings.Replace(rawRules, "\"", "", -1), " ")
	mapRules := make(map[string]string, 12)
	for _, rule := range rules {
		keyValue := strings.Split(rule, ":")
		mapRules[string(keyValue[0])] = string(keyValue[1])
	}

	return mapRules
}