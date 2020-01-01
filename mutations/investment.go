package mutations

import (
	"apijoy/function"
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
	"time"
)

func GetChooseInvestmentMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.ConcurrentResultType,
		Description: "Mutation untuk pemilihan investment",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID investor",
			},
			"project_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID project",
			},
			"quantity": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter jumlah investment",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)
			projectId := p.Args["project_id"].(int)
			quantity := p.Args["quantity"].(int)

			current_time := time.Now()

			// check if user is investor
			// ...

			var period int
			Database.Raw("select period from proposals where id = $1", projectId).Row().Scan(&period)

			investConc := model.Investment_concurrent{
				Investor_id:     investorId,
				Project_id:      projectId,
				Concurrent_code: function.GenerateCode(20),
				Period:          period,
				Quantity:        quantity,
				Status:          "Reviewed",
				Created_at:      current_time,
				Updated_at:      current_time,
				Expire_at:       current_time.Add(time.Hour * time.Duration(2)),
			}

			update := model.Investment_update{
				Project_id: projectId,
				Related_id: investConc.ID,
				Category:   "Funded",
			}

			var result model.UserConcurrent
			if !Database.NewRecord(investConc) && !Database.NewRecord(update) {
				result.Message = "Gagal memilih investment"
				return result, nil
			}
			Database.Create(&investConc)
			Database.Create(&update)

			var amount int
			Database.Raw("select amount from proposals where id = $1", projectId).Row().Scan(&amount)

			transaction := model.Transaction{
				Investor_id: investorId,
				Related_id:  investConc.ID,
				Amount:      amount * quantity,
				Type:        "Send",
				Unique_code: function.GenerateUniqueCode(100, 999),
			}

			Database.Create(&transaction)

			result.Message = "Berhasil memilih investment"
			result.Concurrent = investConc
			result.Sub_total = transaction.Amount
			result.Unique_code = transaction.Unique_code
			result.Total = result.Sub_total + result.Unique_code

			return result, nil
		},
	}
}

func VerifyConcurrentTransaction() *graphql.Field {
	return &graphql.Field{
		Type:        object.VerifiedTransactionType,
		Description: "Memverifikasi transaksi investment concurrent",
		Args: graphql.FieldConfigArgument{
			"transaction_code": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "kode transaksi",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// getting argument values
			code := p.Args["transaction_code"].(string)

			current_time := time.Now()

			var transaction model.Transaction
			Database.Where(&model.Transaction{Transaction_code: code}).First(&transaction)
			transaction.Confirmed_at = current_time
			Database.Save(&transaction)

			var investConc model.Investment_concurrent
			Database.Where(&model.Investment_concurrent{ID: transaction.Related_id, Investor_id: transaction.Investor_id, Status:"Reviewed"}).First(&investConc)
			investConc.Status = "Accepted"
			Database.Save(&investConc)

			var result model.VerifiedTransaction
			result.Message = "Transaction has been verified"
			result.Concurrent = investConc
			result.Transaction = transaction

			return result, nil
		},
	}
}

func GetInvestReleaseMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.InvestmentType,
		Description: "Mutation untuk proses release investment",
		Args: graphql.FieldConfigArgument{
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID investasi",
			},
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investId := p.Args["investment_id"].(int)
			investorId := p.Args["investor_id"].(int)

			// getting current time in format
			current_time := time.Now()

			Database.Exec("update investment_concurrents set release_at = $3 and status = 'Cancelled' "+
				"where investor_id = $1 and id = $2", investorId, investId, current_time)

			return "Investment_concurrent has been released", nil
		},
	}
}

func GetUserInvestAgreementMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.ConcurrentType,
		Description: "Mutation untuk user agreement pada investment concurrent",
		Args: graphql.FieldConfigArgument{
			"concurrent_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID concurrent",
			},
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID investor",
			},
			"is_agree": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "Parameter agreement (true/false)",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			concurrentId := p.Args["concurrent_id"].(int)
			investorId := p.Args["investor_id"].(int)
			isAgree := p.Args["is_agree"].(bool)

			if isAgree {
				Database.Exec("update investment_concurrents set is_agree = 1 "+
					"where id = $1 and investor_id = $2", concurrentId, investorId)
			} else {
				Database.Exec("update investment_concurrents set is_agree = 0 "+
					"where id = $1 and investor_id = $2", concurrentId, investorId)
			}

			var concurrent model.Investment_concurrent
			Database.First(&concurrent, concurrentId)

			return concurrent, nil
		},
	}
}

func GetChangeReinvestMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.InvestmentType,
		Description: "mengubah is_reinvest",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment",
			},
			"is_reinvest": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "status reinvest",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)
			investmentId := p.Args["investment_id"].(int)
			reinvest := p.Args["is_reinvest"].(bool)

			var investment model.Investment
			Database.Where(&model.Investment{ID: investmentId, Investor_id: investorId}).First(&investment)

			if reinvest {
				investment.Is_reinvest = 1
			} else {
				investment.Is_reinvest = 0
			}
			Database.Save(&investment)

			return investment, nil
		},
	}
}

func GetChangeSharedMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.InvestmentReviewType,
		Description: "mengubah is_shared",
		Args: graphql.FieldConfigArgument{
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment",
			},
			"is_shared": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "status shared",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investmentId := p.Args["investment_id"].(int)
			shared := p.Args["is_shared"].(bool)

			var investment model.Investment_review
			Database.Where(&model.Investment_review{Investment_id: investmentId}).First(&investment)

			if shared {
				investment.Is_shared = 1
			} else {
				investment.Is_shared = 0
			}
			Database.Save(&investment)

			return investment, nil
		},
	}
}

func GetInvestmentReviewMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.InvestmentReviewType,
		Description: "menambahkan review investment",
		Args: graphql.FieldConfigArgument{
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment",
			},
			"review": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "komentar review",
			},
			"satisfaction": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "satisfaction level",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investmentId := p.Args["investment_id"].(int)
			review := p.Args["review"].(string)
			satisfaction := p.Args["satisfaction"].(string)

			investment := model.Investment_review{
				Investment_id:      investmentId,
				Review:             review,
				Satisfaction_level: satisfaction,
			}
			Database.Create(&investment)

			return investment, nil
		},
	}
}
