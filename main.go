package main

import (
	"net/http"
	"encoding/json"
	"reflect"
	"go-request-validator/cmd/validator"
)

type test_struc struct {
// 	varName varType `json:"keyName"`
	Name 	string 	`json:"name" type:"string" min:"6" max:"60" required:"true"`
	Email 	string 	`json:"email" type:"email" min:"6" max:"60" required:"true"`
	Age 	int		`json:"age" type:"int" min:"18" max:"90" required:"true"`
}
// Func Handler
func getReq(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var t test_struc
	err := dec.Decode(&t)
	if(err != nil) {
		panic(err)
	}
	validator.Validate(reflect.ValueOf(t))
}

func main()  {
//  http.HandleFunc("/endpoint", func)
	http.HandleFunc("/", getReq)

//  http.ListenAndServe(":port", nil)
	http.ListenAndServe(":8080", nil)
}
