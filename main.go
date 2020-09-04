package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type test_struc struct {
// 	varName varType `json:"keyName"`
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
	Age 	int		`json:"age"`
}

// Func Handler
func getReq(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var t test_struc
	err := dec.Decode(&t)
	if(err != nil) {
		panic(err)
	}
	log.Println(t)
}

func main()  {
//  http.HandleFunc("/endpoint", func)
	http.HandleFunc("/", getReq)

//  http.ListenAndServe(":port", nil)
	http.ListenAndServe(":8080", nil)
}
