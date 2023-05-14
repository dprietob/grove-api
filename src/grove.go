package main

import (
	"fmt"
	"net/http"
	"utils"
)

// type Param struct {
// 	key   string
// 	value string
// }

func main() {
	http.HandleFunc("/", GetRequest)
	http.ListenAndServe(":8080", nil)
}

func GetRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "URI: ", request.RequestURI)
	route, params := utils.DecomposeURI(request.RequestURI)

	fmt.Println(route)
	fmt.Println(params)
}

// func DecomposeURI(uri string) ([]string, []Param) {
// 	var route []string
// 	var params []Param

// 	uriSplitted := strings.Split(uri, "/")

// 	for _, split := range uriSplitted {
// 		if strings.Contains(split, "?") {
// 			subsplit := strings.Split(split, "?")
// 			if len(subsplit) == 2 {
// 				route = append(route, subsplit[0])
// 				params = DecomposeParams(subsplit[1])
// 			}
// 		} else {
// 			route = append(route, split)
// 		}
// 	}

// 	return route, params
// }

// func DecomposeParams(uriParams string) []Param {
// 	var params []Param

// 	paramsSplitted := strings.Split(uriParams, "&")

// 	for _, split := range paramsSplitted {
// 		if strings.Contains(split, "=") {
// 			subsplit := strings.Split(split, "=")
// 			if len(subsplit) == 2 {
// 				params = append(
// 					params,
// 					Param{key: subsplit[0], value: subsplit[1]},
// 				)
// 			}
// 		} else {
// 			params = append(params, Param{key: split, value: "1"})
// 		}
// 	}

// 	return params
// }
