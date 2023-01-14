package main

import (
	"fmt"
	"log"
	"os"
	"thegame/controller"
	"thegame/graph"
	"thegame/middleware"
	"thegame/pkg/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Defining the Graphql handler
func graphqlHandler(database *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &controller.Resolver{}}))
	return middleware.CreateDbContext(database, h)
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	fmt.Println("llsl ", os.Getenv("USER"))
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler(database))
	r.GET("/", playgroundHandler())
	r.Run()
}
