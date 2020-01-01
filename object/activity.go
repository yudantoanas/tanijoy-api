package object

import "github.com/graphql-go/graphql"

var ActivityCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ActivityCategory",
	Description: "Object ActivityCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID activity category",
		},
		"category": &graphql.Field{
			Type:        graphql.String,
			Description: "jenis category",
		},
	},
})
