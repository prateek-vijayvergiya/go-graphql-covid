package Queries

import (
	"Go-Graphql/Types"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

var CovidData = &graphql.Field{
	Type: Types.CovidDataResponseType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		req, err := http.Get("https://data.covid19india.org/data.json")
		if err != nil {
			fmt.Print("error" + err.Error())
		}
		resp, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Print("error" + err.Error())
		}
		var res interface{}
		json.Unmarshal([]byte(resp), &res)
		fmt.Println("res", res)
		return res, nil
	},
}
