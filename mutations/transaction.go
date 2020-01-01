package mutations

import (
	"apijoy/function"
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
	"time"
)

func GetChoosePaymentMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.TransactionResultType,
		Description: "Mutation untuk pemilihan metode pembayaran",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID investor",
			},
			"concurrent_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID concurrent",
			},
			"note": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter note tambahan",
			},
			"payment_method": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter metode pembayaran",
			},
			"payment_account": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "nomor rekening bank",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)
			concurrent := p.Args["concurrent_id"].(int)
			note := p.Args["note"].(string)
			method := p.Args["payment_method"].(string)
			account := p.Args["payment_account"].(string)

			var transaction model.Transaction
			var result model.TransactionResult

			Database.Where(&model.Transaction{Investor_id: investorId, Related_id: concurrent}).First(&transaction)
			transaction.Transaction_code = "Trans-" + function.GenerateCode(5)
			transaction.Note = note
			transaction.Payment_method = method
			transaction.Payment_account = account
			Database.Save(&transaction)
			Database.Exec("update transactions set transaction_at = $2 "+
				"where transaction_code = $1", transaction.Transaction_code, time.Now())

			result.Message = "Berhasil memilih payment method"
			result.Transaction = transaction
			result.Sub_total = transaction.Amount
			result.Unique_code = transaction.Unique_code
			result.Total = result.Sub_total + result.Unique_code
			return result, nil
		},
	}
}

func GetPaymentConfirmMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.TransactionResultType,
		Description: "Mutation untuk konfirmasi pembayaran",
		Args: graphql.FieldConfigArgument{
			"transaction_code": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Kode transaksi",
			},
			"is_confirmed": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "Status konfirm pembayaran (true ketika button konfirm di klik)",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			code := p.Args["transaction_code"].(string)
			confirmed := p.Args["is_confirmed"].(bool)

			var result model.TransactionResult
			var transaction model.Transaction

			if !confirmed {
				result.Message = "Belum konfirmasi"
				return result, nil
			}

			current_time := time.Now()

			// Jika confirmed_at sudah not null maka gabisa diotak atik lagi statusnya
			// .....

			Database.Exec("update transactions set transaction_at = $2 where transaction_code = $1", code, current_time)
			Database.Exec("update investment_concurrents set status = 'Accepted', expire_at = nulll " +
				"where id = (select related_id from transactions where transaction_code = $1)", code)
			Database.Where(&model.Transaction{Transaction_code: code}).First(&transaction)

			result.Message = "Berhasil konfirmasi"
			result.Transaction = transaction
			result.Sub_total = transaction.Amount
			result.Unique_code = transaction.Unique_code
			result.Total = result.Sub_total + result.Unique_code
			return result, nil
		},
	}
}
