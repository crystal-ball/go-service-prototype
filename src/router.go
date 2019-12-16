package main

import "github.com/julienschmidt/httprouter"

// InitRouter sets up the service routes
func InitRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	return router
}
