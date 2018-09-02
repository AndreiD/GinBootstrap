package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/motionwerkGmbH/cpo-backend-api/configs"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-contrib/cors"
	"fmt"
)

var router *gin.Engine

func main() {

	Config := configs.Load()

	if (Config.GetString("environment")) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router = gin.New()
	router.Use(gin.Recovery())

	// allow all origins
	router.Use(cors.Default())

	InitializeRoutes()

	// Serve static invoices files
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))

	// Establish database connection
	//tools.Connect("_theDb.db")
	//tools.MySQLConnect()

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~ SERVER IS RUNNING v: 0.01 ~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// Serve 'em...
	server := &http.Server{
		Addr:           Config.GetString("hostname") + ":" + strconv.Itoa(Config.GetInt("port")),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.SetKeepAlivesEnabled(true)
	server.ListenAndServe()

}
