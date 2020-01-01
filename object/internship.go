package object

import "github.com/graphql-go/graphql"

var InternshipType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Internship",
	Description:"Object Internship untuk menyimpan data registrasi internship",
	Fields: graphql.Fields{
		"id":& graphql.Field{
			Description:"ID dari internship",
			Type:graphql.Int,
		},
		"gpa":& graphql.Field{
			Description:"IPK dari internship",
			Type:graphql.Int,
		},
		"gpa_max":& graphql.Field{
			Description:"IPK tertinggi dari internship",
			Type:graphql.Int,
		},
		"name":& graphql.Field{
			Description:"Nama lengkap dari internship",
			Type:graphql.String,
		},
		"email":& graphql.Field{
			Description:"Email dari internship",
			Type:graphql.String,
		},
		"handphone":& graphql.Field{
			Description:"Nomor handphone dari internship",
			Type:graphql.String,
		},
		"address":& graphql.Field{
			Description:"Alamat lengkap dari internship",
			Type:graphql.String,
		},
		"institute":& graphql.Field{
			Description:"Nama institusi dari internship",
			Type:graphql.String,
		},
		"major":& graphql.Field{
			Description:"Nama jurusan dari internship",
			Type:graphql.String,
		},
	},
})
