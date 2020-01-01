package object

import "github.com/graphql-go/graphql"

var StageType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Stage",
	Description: "Object stage",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID stage",
			Type:        graphql.Int,
		},
		"name": &graphql.Field{
			Description: "Stage name",
			Type:        graphql.String,
		},
		"purpose": &graphql.Field{
			Description: "Stage purpose",
			Type:        graphql.String,
		},
		"description": &graphql.Field{
			Description: "Stage description",
			Type:        graphql.String,
		},
	},
})
