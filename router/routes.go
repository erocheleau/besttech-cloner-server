package router

import (
	"net/http"

	"github.com/erocheleau/besttech-cloner-server/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Install",
		"POST",
		"/install",
		handlers.Install,
	},
	Route{
		"Uninstall",
		"POST",
		"/uninstall",
		handlers.Uninstall,
	},
}
