package queries

import (
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
	"time"
	"strconv"
)

func GetAllInvestmentsQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentType),
		Description: "Query untuk menampilkan list semua project",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var investments []model.Investment
			Database.Find(&investments)

			return investments, nil
		},
	}
}

func GetAllConcurrentsQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.ConcurrentType),
		Description: "Query untuk menampilkan list semua project",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var concurrents []model.Investment_concurrent
			Database.Find(&concurrents)

			return concurrents, nil
		},
	}
}

func GetUserConcurrentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.ConcurrentResultType,
		Description: "Query untuk menampilkan user investment concurrent di halaman pembayaran",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID investor",
			},
			"concurrent_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID concurrent",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			investorId := p.Args["investor_id"].(int)
			concurrentId := p.Args["concurrent_id"].(int)

			var investment model.Investment_concurrent
			var concurrent model.UserConcurrent

			row := Database.Where(&model.Investment_concurrent{ID: concurrentId, Investor_id: investorId}).First(&investment)

			if (row.Error != nil) {
				concurrent.Message = "record not found"
				return concurrent, nil
			}

			var amount, uCode int
			Database.Raw("select amount, unique_code from transactions "+
				"where investor_id = $1 and related_id = $2", investorId, concurrentId).Row().Scan(&amount, &uCode)

			concurrent.Message = "record found"
			concurrent.Concurrent = investment
			concurrent.Sub_total = amount
			concurrent.Unique_code = uCode
			concurrent.Total = concurrent.Sub_total + concurrent.Unique_code

			return concurrent, nil
		},
	}
}

func GetUserInvestmentsQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentType),
		Description: "Query untuk menampilkan list semua proyek invetasi user",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			investorId := p.Args["investor_id"].(int)

			var investments []model.Investment
			Database.Where(&model.Investment{Investor_id: investorId}).Find(&investments)

			return investments, nil
		},
	}
}

func GetUserTotalInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Query untuk menampilkan total investment dari user",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			investorId := p.Args["investor_id"].(int)

			var amount int
			Database.Raw("select sum(amount) from transactions where id in "+
				"(select transaction_id from investments where investor_id = $1);", investorId).Row().Scan(&amount)

			return amount, nil
		},
	}
}

func GetInvestmentPercentageQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Persentase investasi",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var send, receive int
			var result string

			Database.Raw("select sum(amount) from transactions "+
				"where investor_id = $1 and type = 'Receive'", investorId).Row().Scan(&receive)

			Database.Raw("select sum(amount) from transactions "+
				"where investor_id = $1 and type = 'Send'", investorId).Row().Scan(&send)

			if !(receive > send) {
				result = "-" + strconv.FormatFloat((float64(receive)/float64(send))*100, 'f', 0, 64)
				return result, nil
			}
			result = "+" + strconv.FormatFloat((float64(receive)/float64(send))*100, 'f', 0, 64)

			return result, nil
		},
	}
}

func GetUserHasilInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "Hasil investasi investor",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var result int
			Database.Raw("select sum(actual_return) from investments "+
				"where investor_id = $1 "+
				"group by investor_id", investorId).Row().Scan(&result)

			return result, nil
		},
	}
}

func GetUserInvestmentLastUpdatedQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Investment activity dari si user",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var update time.Time
			Database.Raw("select max(updated_at) from transactions "+
				"where investor_id = $1 and type = 'Receive' "+
				"group by investor_id", investorId).Row().Scan(&update)

			result := update.Format("2006-01-02")
			if result == "0001-01-01" {
				return "belum ada hasil investasi", nil
			}

			return result, nil
		},
	}
}

func GetInvestmentActivitiesQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentActivityType),
		Description: "Investment activity dari si user",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// get argument value(s)
			investorId := p.Args["investor_id"].(int)

			var investmentActivities []model.Investment_activity
			rows, _ := Database.Raw("select * from investment_activities "+
				"where investment_id in (select id from investments where investor_id = $1) "+
				"and updated_at > NOW() - INTERVAL '7 days'", investorId).Rows()
			for rows.Next() {
				var activity model.Investment_activity
				Database.ScanRows(rows, &activity)
				investmentActivities = append(investmentActivities, activity)
			}

			return investmentActivities, nil
		},
	}
}

func GetInvestmentActivitiesNotificationQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.ActivitiesNotificationType,
		Description: "angka notifikasi activites",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var notif model.ActivitiesNotification
			Database.Raw("select count(*) from investment_activities "+
				"where investment_id in (select id from investments where investor_id = $1) "+
				"and updated_at > NOW() - INTERVAL '7 days'", investorId).Row().Scan(&notif.Latest_update)

			return notif, nil
		},
	}
}

func GetFeaturedProjectsQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.ProjectType,
		Description: "Proyek rekomendasi",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			var project model.Project
			Database.Where(&model.Project{Is_featured: 1}).First(&project)

			return project, nil
		},
	}
}

func GetRecentInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentType),
		Description: "list investasi baru selesai",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var investments []model.Investment
			rows, _ := Database.Raw("select * from investments where investor_id = $1 "+
				"and actual_finish_at is not null "+
				"and actual_finish_at > now() - interval '7 days'", investorId).Rows()

			for rows.Next() {
				var investment model.Investment
				Database.ScanRows(rows, &investment)
				investments = append(investments, investment)
			}

			return investments, nil
		},
	}
}

func GetHistoryInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentType),
		Description: "list history investasi",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var investments []model.Investment
			rows, _ := Database.Raw("select * from investments where investor_id = $1 "+
				"and actual_finish_at is not null "+
				"and actual_finish_at < now() - interval '7 days'", investorId).Rows()

			for rows.Next() {
				var investment model.Investment
				Database.ScanRows(rows, &investment)
				investments = append(investments, investment)
			}

			return investments, nil
		},
	}
}

func GetOccurringInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.InvestmentType),
		Description: "list investasi berlangsung",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)

			var investments []model.Investment
			rows, _ := Database.Raw("select * from investments where investor_id = $1 "+
				"and actual_finish_at is null", investorId).Rows()

			for rows.Next() {
				var investment model.Investment
				Database.ScanRows(rows, &investment)
				investments = append(investments, investment)
			}

			return investments, nil
		},
	}
}

func GetFinishedInvestmentQueries() *graphql.Field {
	return &graphql.Field{
		Type:        object.InvestmentType,
		Description: "Investment finished",
		Args: graphql.FieldConfigArgument{
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investorId := p.Args["investor_id"].(int)
			investmentId := p.Args["investment_id"].(int)

			var investment model.Investment
			rows, _ := Database.Raw("select * from investments "+
				"where id = $1 "+
				"and investor_id = $2 "+
				"and actual_finish_at is not null", investmentId, investorId).Rows()
			for rows.Next() {
				Database.ScanRows(rows, &investment)
			}

			return investment, nil
		},
	}
}

func GetDetailInvestmentPhotosQueries() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(graphql.String),
		Description: "Detail investasi gallery",
		Args: graphql.FieldConfigArgument{
			"investment_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investment",
			},
			"investor_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "ID investor",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investmentId := p.Args["investment_id"].(int)
			investorId := p.Args["investor_id"].(int)

			var photos []string
			rows, _ := Database.Raw("select src from investment_galleries "+
				"where investment_id = (select id from investments "+
				"where id = $1 and investor_id = $2)", investmentId, investorId).Rows()
			for rows.Next() {
				var src string
				rows.Scan(&src)
				photos = append(photos, src)
			}

			return photos, nil
		},
	}
}

func GetDetailInvestmentUpdateQueries() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(object.InvestmentActivityType),
		Args: graphql.FieldConfigArgument{
			"investment_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"investor_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			investmentId := p.Args["investment_id"].(int)

			var updates []model.Investment_activity
			Database.Where(&model.Investment_activity{Investment_id: investmentId}).Find(&updates)

			return updates, nil
		},
	}
}