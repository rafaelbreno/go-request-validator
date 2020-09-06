package main

import (
	"log"
	"net/http"
	"encoding/json"
	"reflect"
)

type test_struc struct {
// 	varName varType `json:"keyName"`
	Name 	string 	`json:"name" type:"string" min:"6" max:"60" required:"true"`
	Email 	string 	`json:"email" type:"email" min:"6" max:"60" required:"true"`
	Age 	int		`json:"age" type:"int" min:"18" max:"90" required:"true"`
}

func validate(t test_struc) {
	v := reflect.ValueOf(t)

	typeValid := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeValid.Field(i)
		log.Println("Field:", field.Name, " - Value:", v.Field(i))
		log.Println("-_-_-_-_-_-_-_-_-_-_-_-")
	}
}
// Func Handler
func getReq(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var t test_struc
	err := dec.Decode(&t)
	if(err != nil) {
		panic(err)
	}
	validate(t)
}

func main()  {
//  http.HandleFunc("/endpoint", func)
	http.HandleFunc("/", getReq)

//  http.ListenAndServe(":port", nil)
	http.ListenAndServe(":8080", nil)
}
