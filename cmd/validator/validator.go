package validator

import (
	"log"
	"strings"
	"reflect"
	"go-request-validator/cmd/validation"
    "strconv"
)

func Validate(v reflect.Value) {
	typeValid := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeValid.Field(i)
		// log.Println(field.Tag)
		rules := filterRules(string(field.Tag))
		for ruleKey, ruleValue := range rules {
			// log.Println(ruleKey, ruleValue)
			intRuleValue, _ := strconv.Atoi(ruleValue)
			fieldName := v.Field(i).String()
			switch ruleKey {
				case "required":
					log.Println(validation.Required(fieldName, field.Name))
				case "max":
					log.Println(validation.Max(fieldName, field.Name, intRuleValue))
				case "min":
					log.Println(validation.Min(fieldName, field.Name, intRuleValue))
				default: 
					log.Println("-------")
			}
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