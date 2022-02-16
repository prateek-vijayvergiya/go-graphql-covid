package Types

import "github.com/graphql-go/graphql"

var CovidDataFields = graphql.NewObject(graphql.ObjectConfig{
	Name: "Response",
	Fields: graphql.Fields{
		"dailyconfirmed": &graphql.Field{
			Type: graphql.String,
		},
		"dailydeceased": &graphql.Field{
			Type: graphql.String,
		},
		"dailyrecovered": &graphql.Field{
			Type: graphql.String,
		},
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"dateymd": &graphql.Field{
			Type: graphql.String,
		},
		"totalconfirmed": &graphql.Field{
			Type: graphql.String,
		},
		"totaldeceased": &graphql.Field{
			Type: graphql.String,
		},
		"totalrecovered": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CovidDataResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CovidDataResponse",
	Fields: graphql.Fields{
		"cases_time_series": &graphql.Field{
			Type: graphql.NewList(CovidDataFields),
		},
	},
})
