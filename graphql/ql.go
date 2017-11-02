package graphql

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{},
	},
})
