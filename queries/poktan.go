package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
)

func GetAllPoktanQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.PoktanType),
		Description: "Query untuk menampilkan list semua poktan",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var poktans []model.Poktan
			Database.Find(&poktans)

			return poktans, nil
		},
	}
}

func ListUnverifiedPoktan() *graphql.Field {
	return &graphql.Field{
		Type:graphql.NewList(object.UnverifiedPoktan),
		Description: "Query untuk menampilkan poktan yang belum terverifikasi",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var poktans []model.UnverifiedPoktan
			rows, _ := Database.Raw("select name, email, phone, address, vegetables, " +
				"area as area_width, farmers_count as total_farmer, created_at as register_at from poktan_candidates " +
				"where stage_id = 1").Rows()
			for rows.Next() {
				var poktan model.UnverifiedPoktan
				Database.ScanRows(rows, &poktan)

				poktans = append(poktans, poktan)
			}

			return poktans, nil
		},
	}
}
