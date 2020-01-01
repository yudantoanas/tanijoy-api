package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var PaymentBankType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PaymentBank",
	Description: "Objek bank yang digunakan di halaman payment",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Type: graphql.String,
		},
		"logo_path": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BankAccountType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BankAccount",
	Description: "Object Bank Account untuk menyimpan data lengkap rekening bank",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari Bank Account",
			Type:        graphql.Int,
		},
		"bank": &graphql.Field{
			Description: "Nama bank dari Bank Account",
			Type:        graphql.String,
		},
		"alias": &graphql.Field{
			Description: "Alias dari Bank Account",
			Type:        graphql.String,
		},
		"bank_code": &graphql.Field{
			Description: "Kode bank dari Bank Account",
			Type:        graphql.String,
		},
		"account": &graphql.Field{
			Description: "Nomor rekening dari Bank Account",
			Type:        graphql.String,
		},
		"holdername": &graphql.Field{
			Description: "Nama pemilik rekening dari Bank Account",
			Type:        graphql.String,
		},
		"branch": &graphql.Field{
			Description: "Nama cabang dari Bank Account",
			Type:        graphql.String,
		},
		"logo_path": &graphql.Field{
			Description: "logo bank",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Source.(model.Bank_account).Bank

				var path string
				var bank model.Bank_list
				Database.Where(&model.Bank_list{Name: name}).First(&bank)
				path = bank.Logo_path

				return path, nil
			},
		},
	},
})
