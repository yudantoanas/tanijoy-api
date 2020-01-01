package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var ActivitiesNotificationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ActivitiesNotification",
	Description: "angka notifikasi investment activities di halaman investor dashboard",
	Fields: graphql.Fields{
		"latest_update": &graphql.Field{
			Description: "angka notifikasi latest update",
			Type:        graphql.Int,
		},
		"invest_update": &graphql.Field{
			Description: "angka notifikasi investment update",
			Type:        graphql.Int,
		},
	},
})

var InvestmentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Investment",
	Description: "Object Investment untuk menyimpan data investment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID investment",
			Type:        graphql.Int,
		},
		"investor": &graphql.Field{
			Description: "Object investor",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.Investment).Investor_id

				var user model.User
				Database.Find(&user, userId)

				return user, nil
			},
		},
		"project": &graphql.Field{
			Description: "Object project",
			Type:        ProjectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Investment).Project_id

				var project model.Project
				Database.Find(&project, projectId)

				return project, nil
			},
		},
		"land": &graphql.Field{
			Description: "Object land",
			Type:        LandType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				landId := p.Source.(model.Investment).Land_id

				var land model.Land
				Database.First(&land, landId)

				return land, nil
			},
		},
		"transaction": &graphql.Field{
			Description: "Object transaction",
			Type:        TransactionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				transactionId := p.Source.(model.Investment).Transaction_id

				var transaction model.Transaction
				Database.First(&transaction, transactionId)

				return transaction, nil
			},
		},
		"fieldmanager": &graphql.Field{
			Description: "Object fieldmanager",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.Investment).Fieldmanager_id

				var user model.User
				Database.First(&user, userId)

				return user, nil
			},
		},
		"farmer": &graphql.Field{
			Description: "Object farmer",
			Type:        FarmerType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				farmerId := p.Source.(model.Investment).Farmer_id

				var farmer model.Farmer
				Database.First(&farmer, farmerId)

				return farmer, nil
			},
		},
		"stage": &graphql.Field{
			Description: "Object stage",
			Type:        StageType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				stageId := p.Source.(model.Investment).Stage_id

				var stage model.Stage
				Database.First(&stage, stageId)

				return stage, nil
			},
		},
		"phase": &graphql.Field{
			Description: "Object phase",
			Type:        StageType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				phaseId := p.Source.(model.Investment).Phase_id

				var phase model.Stage
				Database.First(&phase, phaseId)

				return phase, nil
			},
		},
		"investment_code": &graphql.Field{
			Description: "Kode investment",
			Type:        graphql.String,
		},
		"is_reinvest": &graphql.Field{
			Description: "Status reinvest (true/false)",
			Type:        graphql.Int,
		},
		"is_reviewed": &graphql.Field{
			Description: "Status reviewed (true/false)",
			Type:        graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.Investment).ID

				var exists bool
				Database.Raw("select exists(select 1 from investment_review "+
					"where investment_id = $1)", id).Row().Scan(&exists)

				return exists, nil
			},
		},
		"delay": &graphql.Field{
			Description: "Lama delay",
			Type:        graphql.Int,
		},
		"sketch_path": &graphql.Field{
			Description: "Path image skecth",
			Type:        graphql.String,
		},
		"certificate_path": &graphql.Field{
			Description: "Path image certificate",
			Type:        graphql.String,
		},
		"closing": &graphql.Field{
			Description: "Investment closing",
			Type:        InvestmentClosingType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.Investment).ID

				var closing model.Investment_closing
				Database.Where(&model.Investment_closing{Investment_id: id}).First(&closing)

				return closing, nil
			},
		},
		"user_review": &graphql.Field{
			Description: "Investment review",
			Type:        InvestmentReviewType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.Investment).ID

				var review model.Investment_review
				Database.Where(&model.Investment_review{Investment_id: id}).First(&review)

				return review, nil
			},
		},
	},
})

var InvestmentClosingType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InvestmentClosing",
	Description: "Object investmentclosing",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"statement": &graphql.Field{
			Type: graphql.String,
		},
		"file_path": &graphql.Field{
			Type: graphql.String,
		},
		"actual_return": &graphql.Field{
			Type: graphql.Int,
		},
		"actual_percentage": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var InvestmentReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InvestmentReview",
	Description: "Object investmentreview",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"satisfaction_level": &graphql.Field{
			Type: graphql.String,
		},
		"review": &graphql.Field{
			Type: graphql.String,
		},
		"is_shared": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				shared := p.Source.(model.Investment_review).Is_shared

				if shared == 1 {
					return true, nil
				}

				return false, nil
			},
		},
	},
})

var InvestmentActivityType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InvestmentActivity",
	Description: "Object InvestmentActivity",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID investment activities",
		},
		"investment": &graphql.Field{
			Type:        InvestmentType,
			Description: "Objek investment",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				investmentId := p.Source.(model.Investment_activity).Investment_id

				var investment model.Investment
				Database.First(&investment, investmentId)

				return investment, nil
			},
		},
		"activity_category": &graphql.Field{
			Type:        ActivityCategoryType,
			Description: "Tahap kategori aktivitas",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				categoryId := p.Source.(model.Investment_activity).Activity_category_id

				var category model.Activity_category
				Database.First(&category, categoryId)

				return category, nil
			},
		},
		"activity_by": &graphql.Field{
			Type:        UserType,
			Description: "User yang melakukan aktivitas",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.Investment_activity).Activity_by

				var user model.User
				Database.First(&user, userId)

				return user, nil
			},
		},
		"activity": &graphql.Field{
			Type:        graphql.String,
			Description: "Jenis aktivitas",
		},
		"photo": &graphql.Field{
			Type:        ListInvestmentPhotoType,
			Description: "Foto investasi",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				investmentId := p.Source.(model.Investment_activity).Investment_id
				activityid := p.Source.(model.Investment_activity).ID

				var listPhoto model.ListInvestmentPhoto
				var count_photo int

				Database.Raw("select count(src) from investment_galleries "+
					"where investment_activity_id = $1 "+
					"and investment_id = $2", activityid, investmentId).Row().Scan(&count_photo)

				if count_photo <= 5 {
					var photos []string
					rows, _ := Database.Raw("select src from investment_galleries "+
						"where investment_activity_id = $1", activityid).Rows()
					for rows.Next() {
						var src string
						rows.Scan(&src)
						photos = append(photos, src)
					}

					listPhoto.Photos = photos
					return listPhoto, nil
				} else if count_photo > 5 {
					var photos []string
					rows, _ := Database.Raw("select src from investment_galleries "+
						"where investment_activity_id = $1 limit 5", activityid).Rows()
					for rows.Next() {
						var src string
						rows.Scan(&src)
						photos = append(photos, src)
					}
					more_photos := count_photo - 5

					listPhoto.Photos = photos
					listPhoto.More_photos = more_photos
					return listPhoto, nil
				}

				return nil, nil
			},
		},
		"file": &graphql.Field{
			Type:        InvestmentFileType,
			Description: "file",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				investmentId := p.Source.(model.Investment_activity).Investment_id
				activityid := p.Source.(model.Investment_activity).ID

				var file model.Investment_file
				Database.Where(&model.Investment_file{Investment_id: investmentId, Investment_activity_id: activityid}).First(&file)

				return file, nil
			},
		},
	},
})

var InvestmentFileType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InvestmentFile",
	Description: "File investasi",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Description: "title file",
			Type:        graphql.String,
		},
		"path": &graphql.Field{
			Description: "path file",
			Type:        graphql.String,
		},
	},
})

var ListInvestmentPhotoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ListInvestmentPhoto",
	Description: "list foto investasi",
	Fields: graphql.Fields{
		"photos": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"more_photos": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var ConcurrentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Concurrent",
	Description: "Object Concurrent untuk menyimpan data investment concurrent",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID concurrent",
			Type:        graphql.Int,
		},
		"investor": &graphql.Field{
			Description: "Data investor",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				investorId := p.Source.(model.Investment_concurrent).Investor_id

				var user model.User
				Database.First(&user, investorId)

				return user, nil
			},
		},
		"project": &graphql.Field{
			Description: "Data proyek yang dipilih",
			Type:        ProjectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectId := p.Source.(model.Investment_concurrent).Project_id

				var project model.Project
				Database.First(&project, projectId)

				return project, nil
			},
		},
		"concurrent_code": &graphql.Field{
			Description: "Kode concurrent",
			Type:        graphql.String,
		},
		"period": &graphql.Field{
			Description: "Lama periode (dalam hari)",
			Type:        graphql.Int,
		},
		"quantity": &graphql.Field{
			Description: "Jumlah kuantitas proyek",
			Type:        graphql.Int,
		},
		"is_agree": &graphql.Field{
			Description: "Status menyetujui",
			Type:        graphql.Int,
		},
		"status": &graphql.Field{
			Description: "Status concurrent",
			Type:        graphql.String,
		},
		"expire_at": &graphql.Field{
			Description: "concurrent time",
			Type:        graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.Investment_concurrent).ID

				var concurrent model.Investment_concurrent
				Database.First(&concurrent, id)

				return concurrent.Expire_at, nil
			},
		},
	},
})

var ConcurrentResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ConcurrentResult",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"concurrent": &graphql.Field{
			Type: ConcurrentType,
		},
		"sub_total": &graphql.Field{
			Type: graphql.Int,
		},
		"unique_code": &graphql.Field{
			Type: graphql.Int,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
