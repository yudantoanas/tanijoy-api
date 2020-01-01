package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
)

func GetAllBanksQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.PaymentBankType),
		Description: "Query untuk menampilkan list bank pada halaman pembayaran",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rows, _ := Database.Raw("select bank_accounts.account, logo_path from bank_accounts " +
				"left join bank_list on bank_list.alias = bank_accounts.alias").Rows()

			var banks []model.Payment_bank
			for rows.Next() {
				var bank model.Payment_bank
				// scan every rows and save it to bank
				Database.ScanRows(rows, &bank)

				// append each bank to banks array
				banks = append(banks, bank)
			}

			return banks, nil
		},
	}
}

func GetBankQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.BankAccountType,
		Description: "Query untuk menampilkan detail bank pada halaman konfirmasi",
		Args: graphql.FieldConfigArgument{
			"account": &graphql.ArgumentConfig{
				Description: "account number",
				Type:        graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			account := p.Args["account"].(string)

			var bankAccount model.Bank_account
			Database.Where(&model.Bank_account{Account: account}).First(&bankAccount)

			return bankAccount, nil
		},
	}
}
