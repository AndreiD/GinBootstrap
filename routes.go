package main

import "github.com/motionwerkGmbH/cpo-backend-api/handlers"

func InitializeRoutes() {

	v1 := router.Group("/api/v1")
	{
		//~~~~~~~~~~~~~~~~~~~~ GENERAL STUFF ~~~~~~~~~~~~~~~~~~~~~~

		v1.GET("/", handlers.Index)

	}

}
