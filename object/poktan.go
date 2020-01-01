package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var UnverifiedPoktan = graphql.NewObject(graphql.ObjectConfig{
	Name: "UnverifiedPoktan",
	Description:"poktan yang belum terverifikasi",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Description: "poktan name",
			Type:        graphql.String,
		},
		"email": &graphql.Field{
			Description: "poktan email",
			Type:        graphql.String,
		},
		"phone": &graphql.Field{
			Description: "poktan phone",
			Type:        graphql.String,
		},
		"address": &graphql.Field{
			Description: "poktan address",
			Type:        graphql.String,
		},
		"vegetables": &graphql.Field{
			Description: "poktan vegetables",
			Type:        graphql.String,
		},
		"area_width": &graphql.Field{
			Description: "area width",
			Type:        graphql.Int,
		},
		"total_farmer": &graphql.Field{
			Description: "total farmer",
			Type:        graphql.Int,
		},
		"register_at": &graphql.Field{
			Description: "time of registration",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				timeStr := p.Source.(model.UnverifiedPoktan).Register_at

				return timeStr.Format("02 Jan 2006 15:04"), nil
			},
		},
	},
})

var PoktanType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Poktan",
	Description:"Object Poktan untuk menyimpan data kelompok tani tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description:"ID dari kelompok tani",
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Description:"Nama dari kelompok tani",
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Description:"Email dari kelompok tani",
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Description:"Nomor telepon dari kelompok tani",
			Type: graphql.String,
		},
		"area": &graphql.Field{
			Description:"Area dari kelompok tani",
			Type: graphql.String,
		},
		"farmers_count": &graphql.Field{
			Description:"Jumlah petani dari kelompok tani",
			Type: graphql.Int,
		},
		"vegetables": &graphql.Field{
			Description:"Nama sayuran dari kelompok tani",
			Type: graphql.String,
		},
		"stage": &graphql.Field{
			Description:"Stage dari kelompok tani",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				type PoktanStage struct {
					Name string
				}

				var poktanStage PoktanStage

				poktanId := p.Source.(model.Poktan).ID

				Database.Raw("SELECT name FROM poktan_stages "+
					"WHERE id=(SELECT stage_id FROM poktan_candidates WHERE id=$1)", poktanId).Scan(&poktanStage)

				return poktanStage.Name, nil
			},
		},
	},
})
