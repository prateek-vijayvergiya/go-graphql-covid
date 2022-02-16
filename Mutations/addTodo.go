package Mutations

import (
	"Go-Graphql/Types"

	"github.com/graphql-go/graphql"
)

var AddTodo = graphql.NewObject(graphql.ObjectConfig{
	Name: "AddTodo",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type: Types.TodoType,
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				text, _ := p.Args["text"].(string)
				newTodo := &Types.Todo{
					Text: text,
				}

				return newTodo, nil
			},
		},
	},
})
