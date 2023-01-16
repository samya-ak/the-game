package main

import (
	"log"
	"thegame/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/query", graphqlHandler(database))
	r.GET("/", playgroundHandler())
	r.Run()
}
