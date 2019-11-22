package main

import "changeme/handlers"

func InitializeRouter() {

	api := router.Group("/api")

	// pong
	api.GET("/ping", handlers.Ping)

}
