package object

import "github.com/graphql-go/graphql"

// create role object
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Description:"Object Role untuk menyimpan data role user tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description:"ID dari Role",
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Description:"Nama dari Role",
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Description:"Deskripsi dari Role",
			Type: graphql.String,
		},
	},
})

