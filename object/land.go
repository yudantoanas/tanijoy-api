package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var LandType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Land",
	Description: "Object land",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID land",
		},
		"regional": &graphql.Field{
			Type:        RegionalType,
			Description: "Object regional",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				regionalId := p.Source.(model.Land).Regional_id

				var regional model.Regional
				Database.First(&regional, regionalId)

				return regional, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Land name",
		},
		"area": &graphql.Field{
			Type:        graphql.Int,
			Description: "Land area",
		},
		"province": &graphql.Field{
			Type:        graphql.String,
			Description: "Province name ",
		},
		"city": &graphql.Field{
			Type:        graphql.String,
			Description: "City name",
		},
		"address": &graphql.Field{
			Type:        graphql.String,
			Description: "Land address",
		},
		"postal_code": &graphql.Field{
			Type:        graphql.String,
			Description: "Postal code address",
		},
		"ownership": &graphql.Field{
			Type:        graphql.String,
			Description: "Land ownership",
		},
		"certificate": &graphql.Field{
			Type:        graphql.String,
			Description: "Land certificate",
		},
	},
})
