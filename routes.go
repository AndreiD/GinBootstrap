package main

import "github.com/AndreiD/GinBootstrap/handlers"

func InitializeRoutes() {

	v1 := router.Group("/api/v1")
	{
		//~~~~~~~~~~~~~~~~~~~~ GENERAL STUFF ~~~~~~~~~~~~~~~~~~~~~~

		v1.GET("/", handlers.Index)

	}

}
