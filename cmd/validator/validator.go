package validator

import (
	// "fmt"
	"log"
	// "reflect"
	// "go-request-validator/cmd/validation"
	// "strconv"
	"net/http"
	"encoding/json"
	"reflect"
	"strings"
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
	/* need to add the Elem()
	 * because it's a pointer
	*/
	validateRequest(reflect.ValueOf(reqStruct).Elem())
}

func validateRequest(requestValues reflect.Value) {
	structFields := requestValues.Type()
	// log.Println(requestValues.NumField());
	for i := 0; i < requestValues.NumField(); i++ {
		// {Name  string json:"name" type:"string" min:"6" max:"60" required:"true" 0 [0] false}
		field := structFields.Field(i)
		// Field Name
		// log.Println(field.Name)
		
		// map[json:name max:60 min:6 required:true type:string]
		tagRules := filterRules(string(field.Tag))
		for ruleKey, ruleValue := range tagRules {
			// converting rule value to int
			// intRuleValue, _ := strconv.Atoi(ruleValue)
			// fieldRuleName := requestValues.Field(i).String()
			// log.Println(fieldRuleName, intRuleValue)
			log.Println(ruleKey, ruleValue)
		}
	}
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

func filterRules(rawRules string) map[string]string {
	rules := strings.Split(strings.Replace(rawRules, "\"", "", -1), " ")
	mapRules := make(map[string]string, 12)
	for _, rule := range rules {
		keyValue := strings.Split(rule, ":")
		mapRules[string(keyValue[0])] = string(keyValue[1])
	}

	return mapRules
}