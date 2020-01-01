package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var RegionalType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Regional",
	Description: "Object regional",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID regional",
		},
		"manager": &graphql.Field{
			Type:        UserType,
			Description: "Object manager",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.Regional).Manager_id

				var user model.User
				Database.First(&user, userId)

				return user, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.Int,
			Description: "Regional name",
		},
	},
})
