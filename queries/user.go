package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
)

func GetAllUserQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.UserType),
		Description: "Query untuk menampilkan list semua user",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []model.User
			Database.Find(&users)

			return users, nil
		},
	}
}

func GetUserDetailQueries() *graphql.Field {
	return &graphql.Field{
		Type: object.UserType,
		Args: graphql.FieldConfigArgument{
			"userId": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter userId",
			},
		},
		Description: "Query untuk menampilkan detail user",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			userId := p.Args["userId"].(int)

			var user model.User
			Database.First(&user, userId)

			return user, nil
		},
	}
}

func GetPayMethodsQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(graphql.String),
		Description: "Query untuk menampilkan pilihan payment method",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// array of methods
			payMethods := [3]string{"Method 1", "Method 2", "Method 3"}

			return payMethods, nil
		},
	}
}

func GetProfilePercentageQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserInformationType,
		Description: "Query untuk progress profil",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "user id",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)

			var user model.User_information
			Database.Where(&model.User_information{User_id: userId}).First(&user)

			return user, nil
		},
	}
}

func OtherTasks() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "query tugas lain",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return "berganti", nil
		},
	}
}