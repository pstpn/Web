package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"course/internal/controller/v2/graphql/graph"
	"course/internal/service"
)

// Defining the Graphql handler
func graphqlHandler(checkpointService service.CheckpointService) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CheckpointService: checkpointService,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/api/v2/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Handle(handler *gin.Engine, checkpointService service.CheckpointService) {
	handler.POST("graphql/api/v2/query", graphqlHandler(checkpointService))
	handler.GET("graphql/api/v2/", playgroundHandler())
}
