package main

import (
	"Go-Graphql/Mutations"
	"Go-Graphql/Queries"
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"greet":     Queries.Post,
		"covidData": Queries.CovidData,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: Mutations.AddTodo,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	corsHandler := cors.Default().Handler(h)
	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", corsHandler)
	fmt.Println("Server running")
}
