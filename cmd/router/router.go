package router

import (
	"fmt"
	"net/http"
)

func Route(method string, url string) {
	fmt.Println(method, url)
}