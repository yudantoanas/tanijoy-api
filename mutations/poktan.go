package mutations

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
	"time"
)

func GetSignupPoktanMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk proses sign up poktan",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nama poktan",
			},
			"email": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter email poktan",
			},
			"phone": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nomor telepon",
			},
			"address": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter alamat poktan",
			},
			"area": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter area poktan",
			},
			"farmers": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter jumlah petani",
			},
			"vegetables": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nama sayuran",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// generate current time
			current_time := time.Now()
			current_time.Format("2006-01-02 15:04:05")

			newPoktan := model.Poktan{
				Name:          p.Args["name"].(string),
				Email:         p.Args["email"].(string),
				Phone:         p.Args["phone"].(string),
				Address:       p.Args["address"].(string),
				Area:          p.Args["area"].(int),
				Farmers_count: p.Args["farmers"].(int),
				Vegetables:    p.Args["vegetables"].(string),
				Stage_id:      1,
				Created_at:    current_time,
				Updated_at:    current_time,
			}

			var poktan model.Poktan
			Database.Where(&model.Poktan{Email: newPoktan.Email}).First(&poktan)
			if poktan.Email == newPoktan.Email {
				return "Poktan email already exist", nil
			}

			Database.Create(&newPoktan)

			return "Success create poktan", nil
		},
	}
}

func GetAccPoktanMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk proses acc poktan",
		Args: graphql.FieldConfigArgument{
			"poktanId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"is_confirmed": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			poktanId := p.Args["poktanId"].(int)
			isConfirmed := p.Args["is_confirmed"].(bool)

			// Finding poktan with poktanId
			var poktan model.Poktan
			if find := Database.Find(&poktan, poktanId); find.Error != nil {
				// if poktan doesn't exist
				return "Poktan doesn't exist", find.Error
			}

			// Check if poktan is confirmed or not
			if !isConfirmed {
				// if poktan is rejected
				Database.Model(&poktan).Updates(model.Poktan{Stage_id: 5})
				return "Poktan has been rejected", nil
			}

			// poktan is confirmed
			Database.Model(&poktan).Updates(model.Poktan{Stage_id: 4})

			return "Poktan has been accepted", nil
		},
	}
}
