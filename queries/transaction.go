package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
)

func GetConcurrentTimerQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.DateTime,
		Description: "Query untuk waktu expired concurrent pada menu konfirmasi",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
			"concurrent_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment concurrent",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			investorId := p.Args["investor_id"].(int)
			concurrentId := p.Args["concurrent_id"].(int)

			var concurrent model.Investment_concurrent
			Database.Where(&model.Investment_concurrent{ID: concurrentId, Investor_id: investorId}).First(&concurrent)

			return concurrent.Expire_at, nil
		},
	}
}

func UnconfirmedTransaction() *graphql.Field {
	return &graphql.Field{
		Type:        object.TransactionResultType,
		Description: "Query untuk data transaksi pada menu konfirmasi",
		Args: graphql.FieldConfigArgument{
			"transaction_code": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Kode transaksi",
			},
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
			"concurrent_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment concurrent",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			code := p.Args["transaction_code"].(string)
			investorId := p.Args["investor_id"].(int)
			concurrentId := p.Args["concurrent_id"].(int)

			var transaction model.Transaction
			var result model.TransactionResult
			row := Database.Where(&model.Transaction{Transaction_code: code,
				Investor_id: investorId, Related_id: concurrentId}).First(&transaction)

			if (row.Error != nil) {
				result.Message = "record not found"
				return result, nil
			}

			result.Message = "record found"
			result.Transaction = transaction
			result.Sub_total = transaction.Amount
			result.Unique_code = transaction.Unique_code
			result.Total = result.Sub_total + result.Unique_code

			return result, nil
		},
	}
}

func GetTotalTransactionQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Total transaksi",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var total int
			Database.Raw("select count(id) from transactions "+
				"where investor_id = $1 and confirmed_at is not null", investorId).Row().Scan(&total)

			return total, nil
		},
	}
}

func GetCancelledTransactionQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.ConcurrentType),
		Description: "list investment concurrent yang tertunda",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var concurrents []model.Investment_concurrent
			Database.Where(&model.Investment_concurrent{Investor_id: investorId, Status: "Cancelled"}).Find(&concurrents)

			return concurrents, nil
		},
	}
}

func GetUnpaidTransactionQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.TransactionHistory),
		Description: "list concurrent yang belum dibayar",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var histories []model.TransactionHistory
			rows, _ := Database.Table("investment_concurrents").
				Select("investment_concurrents.id as concurrent_id, " +
					"investment_concurrents.investor_id as investor_id, " +
					"investment_concurrents.project_id as project_id, " +
					"investment_concurrents.quantity   as quantity, " +
					"investment_concurrents.is_agree   as is_agree, " +
					"investment_concurrents.status     as status, " +
					"transactions.amount               as amount, " +
					"transactions.transaction_at       as transaction_at").
				Joins("left join transactions " +
					"on transactions.related_id = investment_concurrents.id " +
					"and transactions.investor_id = investment_concurrents.investor_id").
				Where("is_agree = 1 "+
					"and status = 'Cancelled' "+
					"and investment_concurrents.investor_id = $1", investorId).Rows()

			for rows.Next() {
				var result model.TransactionHistory
				var res model.ConcurrentTransaction
				Database.ScanRows(rows, &res)
				result.Concurrent_transaction = res
				result.State = "Payment Section"
				histories = append(histories, result)
			}

			return histories, nil
		},
	}
}

func GetUnagreedTransactionQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.TransactionHistory),
		Description: "list concurrent yang belum di-agree",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var histories []model.TransactionHistory
			rows, _ := Database.Table("investment_concurrents").
				Select("id as concurrent_id, " +
					"investment_concurrents.investor_id as investor_id, " +
					"investment_concurrents.project_id as project_id, " +
					"investment_concurrents.quantity   as quantity, " +
					"investment_concurrents.is_agree   as is_agree, " +
					"investment_concurrents.status     as status").
				Where("is_agree = 0 "+
					"and status = 'Cancelled'"+
					"and investor_id = $1", investorId).Rows()

			for rows.Next() {
				var result model.TransactionHistory
				var res model.ConcurrentTransaction
				Database.ScanRows(rows, &res)
				result.Concurrent_transaction = res
				result.State = "Agreement Section"
				histories = append(histories, result)
			}

			return histories, nil
		},
	}
}

func GetExpiredHistoryQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.TransactionHistory),
		Description: "list riwayat transaksi expired",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var histories []model.TransactionHistory
			rows, _ := Database.Table("investment_concurrents").
				Select("investment_concurrents.id as concurrent_id, " +
					"investment_concurrents.investor_id as investor_id, " +
					"investment_concurrents.project_id as project_id, " +
					"investment_concurrents.quantity   as quantity, " +
					"investment_concurrents.is_agree   as is_agree, " +
					"investment_concurrents.status     as status, " +
					"transactions.amount               as amount, " +
					"transactions.transaction_at       as transaction_at").
				Joins("left join transactions " +
					"on transactions.related_id = investment_concurrents.id " +
					"and transactions.investor_id = investment_concurrents.investor_id").
				Where("investment_concurrents.investor_id = $1 "+
					"and expire_at is not null", investorId).Rows()

			for rows.Next() {
				var result model.TransactionHistory
				var res model.ConcurrentTransaction
				Database.ScanRows(rows, &res)
				result.Concurrent_transaction = res
				result.State = "Expired"
				histories = append(histories, result)
			}

			return histories, nil
		},
	}
}

func GetSuccessHistoryQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.TransactionHistory),
		Description: "list riwayat transaksi berhasil",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var histories []model.TransactionHistory
			rows, _ := Database.Table("investment_concurrents").
				Select("investment_concurrents.id as concurrent_id, " +
					"investment_concurrents.investor_id as investor_id, " +
					"investment_concurrents.project_id as project_id, " +
					"investment_concurrents.quantity   as quantity, " +
					"investment_concurrents.is_agree   as is_agree, " +
					"investment_concurrents.status     as status, " +
					"transactions.amount               as amount, " +
					"transactions.transaction_at       as transaction_at").
				Joins("left join transactions " +
					"on transactions.related_id = investment_concurrents.id " +
					"and transactions.investor_id = investment_concurrents.investor_id").
				Where("investment_concurrents.investor_id = $1 "+
					"and expire_at is null", investorId).Rows()

			for rows.Next() {
				var result model.TransactionHistory
				var res model.ConcurrentTransaction
				Database.ScanRows(rows, &res)
				result.Concurrent_transaction = res
				result.State = "Success"
				histories = append(histories, result)
			}

			return histories, nil
		},
	}
}

func UnverifiedTransaction() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.UnverifiedTransactionType),
		Description: "list transaksi yang belum diverifikasi",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var list []model.UnverifiedTransaction

			rows, _ := Database.Raw("select transaction_code as code, investor_id, " +
				"(select project_id from investment_concurrents where id = transactions.related_id) as project_id, " +
				"(amount + unique_code) as total_amount, " +
				"(select bank from bank_accounts where account = transactions.payment_account) as bank_name, " +
				"transaction_at as time " +
				"from transactions where confirmed_at is null").Rows()
			for rows.Next() {
				var unverif model.UnverifiedTransaction
				Database.ScanRows(rows, &unverif)

				list = append(list, unverif)
			}

			return list, nil
		},
	}
}
