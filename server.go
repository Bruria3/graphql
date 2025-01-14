package main

import (
	"go-react-graphql-orders/middleware"
	"go-react-graphql-orders/resolver"
	"go-react-graphql-orders/schema"
	"go-react-graphql-orders/service"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"

	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	ctx := context.Background()
	orderService := service.NewOrderService()
	ctx = context.WithValue(ctx, service.KeyOrderService, orderService)
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})
	http.Handle("/api/query", corsHandler.Handler(middleware.AddContext(ctx, &middleware.GraphQL{Schema: graphqlSchema})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
