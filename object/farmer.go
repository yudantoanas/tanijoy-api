package object

import "github.com/graphql-go/graphql"

var FarmerType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Farmer",
	Description: "Object Farmer yang menyimpan data petani Tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari petani",
			Type:        graphql.Int,
		},
		"account_id": &graphql.Field{
			Description: "ID akun dari petani",
			Type:        graphql.Int,
		},
		"name": &graphql.Field{
			Description: "Nama dari petani",
			Type:        graphql.String,
		},
		"phone": &graphql.Field{
			Description: "Nomor telepon dari petani",
			Type:        graphql.String,
		},
		"address": &graphql.Field{
			Description: "Alamat dari petani",
			Type:        graphql.String,
		},
		"age": &graphql.Field{
			Description: "Umur dari petani",
			Type:        graphql.Int,
		},
		"year_of_experiences": &graphql.Field{
			Description: "Tahun pengalaman dari petani",
			Type:        graphql.Int,
		},
		"is_married": &graphql.Field{
			Description: "Status pengalaman dari petani",
			Type:        graphql.Int,
		},
		"number_of_childs": &graphql.Field{
			Description: "Jumlah anak dari petani",
			Type:        graphql.Int,
		},
		"specializations": &graphql.Field{
			Description: "Spesialisasi dari petani",
			Type:        graphql.String,
		},
		"profpic_src": &graphql.Field{
			Description: "Profile picture dari petani",
			Type:        graphql.String,
		},
	},
})