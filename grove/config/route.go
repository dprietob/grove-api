package config

import (
	"net/http"
	"regexp"
	"strings"
)

type Route struct {
	method  string
	regexp  string
	handler string
}

type Param struct {
	key   string
	value string
}

func GetRoutes() []Route {
	return []Route{
		{method: "POST", regexp: "/product/save", handler: "product:save"},
		{method: "GET", regexp: "/product/{id:[0-9]+}", handler: "product:view"},
		{method: "GET", regexp: "/product/{id:[0-9]+}/edit", handler: "product:edit"},
		{method: "POST", regexp: "/product/{id:[0-9]+}/edit", handler: "product:persist"},
	}
}

func GetRouteFromURI(request *http.Request) (Route, [][]Param, bool) {
	routes := GetRoutes()
	splUri, getParams := DecomposeURI(request.RequestURI)
	method := request.Method
	var params [][]Param

	for _, route := range routes {
		var uriParams []Param
		splRoute, _ := DecomposeURI(route.regexp)
		good := true

		if len(splUri) != len(splRoute) {
			continue
		}

		for i := 0; i < len(splRoute); i++ {
			key, regExp, isRegexp := GetRegexp(splRoute[i])

			if isRegexp {
				match, _ := regexp.MatchString(regExp, splUri[i])

				if match {
					uriParams = append(
						uriParams,
						Param{key: key, value: splUri[i]},
					)
				} else {
					good = false
					break
				}
			} else {
				if splRoute[i] != splUri[i] {
					good = false
					break
				}
			}
		}

		if good && method == route.method {
			params = append(params, uriParams)
			params = append(params, getParams)

			return route, params, false
		}
	}

	return Route{}, [][]Param{}, true
}

func DecomposeURI(uri string) ([]string, []Param) {
	var route []string
	var params []Param

	uriSplitted := strings.Split(uri, "/")

	for _, split := range uriSplitted {
		if strings.Contains(split, "?") {
			subsplit := strings.Split(split, "?")
			if len(subsplit) == 2 {
				route = append(route, subsplit[0])
				params = DecomposeParams(subsplit[1])
			}
		} else {
			route = append(route, split)
		}
	}

	return route, params
}

func DecomposeParams(uriParams string) []Param {
	var params []Param

	paramsSplitted := strings.Split(uriParams, "&")

	for _, split := range paramsSplitted {
		if strings.Contains(split, "=") {
			subsplit := strings.Split(split, "=")
			if len(subsplit) == 2 {
				params = append(
					params,
					Param{key: subsplit[0], value: subsplit[1]},
				)
			}
		} else {
			params = append(params, Param{key: split, value: "1"})
		}
	}

	return params
}

func GetRegexp(exp string) (string, string, bool) {
	isRegexp := strings.Contains(exp, "{") && strings.Contains(exp, "}")
	key := ""
	regexp := ""

	if isRegexp {
		spl := strings.Split(exp, ":")
		if len(spl) == 2 {
			key = strings.Replace(spl[0], "{", "", -1)
			regexp = strings.Replace(spl[1], "}", "", -1)
		} else {
			isRegexp = false
		}
	}

	return key, regexp, isRegexp
}
