package object

import "github.com/graphql-go/graphql"

var JobPositionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "jobPosition",
	Description:"Object jobPosition untuk menyimpan data posisi pekerjaan",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description:"ID dari posisi pekerjaan",
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Description:"Nama dari posisi pekerjaan",
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Description:"Deskripsi dari posisi pekerjaan",
			Type: graphql.String,
		},
	},
})