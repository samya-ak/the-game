package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"thegame/controller"
	"thegame/graph"
	"thegame/middleware"
	"thegame/model"
)

// Defining the Graphql handler
func graphqlHandler(database *gorm.DB) gin.HandlerFunc {
	c := graph.Config{
		Resolvers: &controller.Resolver{
			UserService: &model.User{},
		},
	}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	return middleware.CreateDbContext(database, h)
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
