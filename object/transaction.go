package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
)

var TransactionResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TransactionResult",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Description: "Message response",
			Type:        graphql.String,
		},
		"transaction": &graphql.Field{
			Description: "Object transaksi",
			Type:        TransactionType,
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

var ConcurrentTransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ConcurrentTransaction",
	Fields: graphql.Fields{
		"concurrent_id": &graphql.Field{
			Type: graphql.Int,
		},
		"investor": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.ConcurrentTransaction).Investor_id

				var investor model.User
				Database.First(&investor, id)

				return investor, nil
			},
		},
		"project": &graphql.Field{
			Type: ProjectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.ConcurrentTransaction).Project_id

				var project model.Project
				Database.First(&project, id)

				return project, nil
			},
		},
		"quantity": &graphql.Field{
			Type: graphql.Int,
		},
		"is_agree": &graphql.Field{
			Type: graphql.Int,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"transaction_at": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				datetime := p.Source.(model.ConcurrentTransaction).Transaction_at

				return datetime.Format("2006-01-02"), nil
			},
		},
	},
})

var TransactionHistory = graphql.NewObject(graphql.ObjectConfig{
	Name: "TransactionHistory",
	Fields: graphql.Fields{
		"concurrent_transaction": &graphql.Field{
			Description: "objek transaksi concurrent",
			Type:        ConcurrentTransactionType,
		},
		"state": &graphql.Field{
			Description: "state riwayat transaksi",
			Type:        graphql.String,
		},
	},
})

var UnverifiedTransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UnverifiedTransaction",
	Description: "Object Transaction untuk menyimpan data transaksi investment tanijoy",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Description: "kode transaksi",
			Type:        graphql.String,
		},
		"investor": &graphql.Field{
			Description: "investor",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.UnverifiedTransaction).Investor_id

				var user model.User
				Database.First(&user, id)

				return user, nil
			},
		},
		"proyek_id": &graphql.Field{
			Description: "proyek",
			Type:        ProjectType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.UnverifiedTransaction).Project_id

				var project model.Project
				Database.First(&project, id)

				return project, nil
			},
		},
		"total_amount": &graphql.Field{
			Description: "total pembayaran",
			Type:        graphql.Int,
		},
		"bank_name": &graphql.Field{
			Description: "nama bank",
			Type:        graphql.String,
		},
		"time": &graphql.Field{
			Description: "waktu transaksi",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				datetime := p.Source.(model.UnverifiedTransaction).Time

				return datetime.Format("02 Jan 2006 15:04"), nil
			},
		},
	},
})

var VerifiedTransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "VerifiedTransaction",
	Description: "Hasil transaksi yang di verified",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"concurrent": &graphql.Field{
			Type: ConcurrentType,
		},
		"transaction": &graphql.Field{
			Type: TransactionType,
		},
	},
})

var TransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Transaction",
	Description: "Object Transaction untuk menyimpan data transaksi investment tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "ID dari Transaksi",
			Type:        graphql.Int,
		},
		"transaction_code": &graphql.Field{
			Description: "Kode dari Transaksi",
			Type:        graphql.String,
		},
		"investor": &graphql.Field{
			Description: "Data investor dari Transaksi",
			Type:        UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				investorId := p.Source.(model.Transaction).Investor_id

				var user model.User
				Database.First(&user, investorId)

				return user, nil
			},
		},
		"concurrent_related_id": &graphql.Field{
			Description: "ID relasi ke investment_concurrent",
			Type:        ConcurrentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				concurrentId := p.Source.(model.Transaction).Related_id

				var concurrent model.Investment_concurrent
				Database.First(&concurrent, concurrentId)

				return concurrent, nil
			},
		},
		"amount": &graphql.Field{
			Description: "Jumlah investasi dalam rupiah",
			Type:        graphql.Int,
		},
		"unique_code": &graphql.Field{
			Description: "Kode unik untuk setiap transaksi yang akan ditambahkan ke ammount",
			Type:        graphql.Int,
		},
		"type": &graphql.Field{
			Description: "Tipe transaksi",
			Type:        graphql.String,
		},
		"note": &graphql.Field{
			Description: "Note tambahan pada transaksi",
			Type:        graphql.String,
		},
		"payment_account": &graphql.Field{
			Description: "Nama rekening bank",
			Type:        graphql.String,
		},
		"payment_method": &graphql.Field{
			Description: "Metode pembayaran",
			Type:        graphql.String,
		},
	},
})
