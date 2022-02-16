package Types

import "github.com/graphql-go/graphql"

type Todo struct {
	Text string `json:"text"`
}

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"text": &graphql.Field{
			Type: graphql.String,
		},
	},
})
