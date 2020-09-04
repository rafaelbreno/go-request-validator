package router

import (
	"fmt"
)

func Route(method string, url string) {
	fmt.Println(method, url)
}