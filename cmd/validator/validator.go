package validator

import (
	// "fmt"
	"log"
	// "strings"
	// "reflect"
	// "go-request-validator/cmd/validation"
	// "strconv"
	"net/http"
	"encoding/json"
	"reflect"
)

/* 
 * Function Validate receives
 * requestValues an instance of Value(reflect package)
 * a
 **/
func Validate(reqStruct interface{}, req *http.Request) {
	// log.Printf("%v", reqStruct)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(reqStruct)

	if(err != nil) {
		panic(err)
	}

	log.Println(reflect.ValueOf(reqStruct))
}

// func Validate(v reflect.Value) {
// 	typeValid := v.Type()

// 	for i := 0; i < v.NumField(); i++ {
// 		field := typeValid.Field(i)
// 		rules := filterRules(string(field.Tag))
// 		for ruleKey, ruleValue := range rules {
// 			intRuleValue, _ := strconv.Atoi(ruleValue)
// 			fieldName := v.Field(i).String()
// 			fieldType := fmt.Sprint(field.Type)
// 			log.Println("field:", field)
// 			log.Println("rules:", rules)
// 			log.Println("ruleValue:", ruleValue)
// 			log.Println("ruleKey:", ruleKey)
// 			log.Println("intRuleValue:", intRuleValue)
// 			log.Println("fieldName:", fieldName)
// 			log.Println("fieldType:", fieldType)
// 			switch ruleKey {
// 				case "required":
// 					log.Println(validation.Required(fieldName, field.Name))
// 					break
// 				case "max":
// 					log.Println(validation.Max(fieldName, field.Name, intRuleValue, fieldType))
// 					break
// 				case "min":
// 					log.Println(validation.Min(fieldName, field.Name, intRuleValue, fieldType))
// 					break
// 				default:
// 					break
// 			}
// 		}
// 	}
// }

// func filterRules(rawRules string) map[string]string {
// 	rules := strings.Split(strings.Replace(rawRules, "\"", "", -1), " ")
// 	mapRules := make(map[string]string, 12)
// 	for _, rule := range rules {
// 		keyValue := strings.Split(rule, ":")
// 		mapRules[string(keyValue[0])] = string(keyValue[1])
// 	}

// 	return mapRules
// }