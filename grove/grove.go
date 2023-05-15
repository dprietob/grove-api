package main

import (
	"fmt"
	"grove/config"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetRequest)
	http.ListenAndServe(":8080", nil)
}

func GetRequest(writer http.ResponseWriter, request *http.Request) {
	route, params, err := config.GetRouteFromURI(request)

	fmt.Println(route)
	fmt.Println(params)
	fmt.Println(err)
	fmt.Println()
}
