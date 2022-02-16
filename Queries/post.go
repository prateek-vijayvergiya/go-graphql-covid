package Queries

import "github.com/graphql-go/graphql"

var Post = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "Hello World!", nil
	},
}
