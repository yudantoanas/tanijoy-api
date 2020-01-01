package object

import (
	"apijoy/model"
	"github.com/graphql-go/graphql"
	"log"
)

var UserResult = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserResult",
	Description: "UserResult menyimpan message dan data User ketika dia telah login/signup",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type:        graphql.String,
			Description: "response message",
		},
		"user": &graphql.Field{
			Type:        UserType,
			Description: "objeck user",
		},
	},
})

var UserInformationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserInformation",
	Description: "Object User information untuk menyimpan data lengkap user tanijoy",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID dari User",
		},
		"identities_number": &graphql.Field{
			Type:        graphql.String,
			Description: "nomor identitas User",
		},
		"province": &graphql.Field{
			Type:        graphql.String,
			Description: "nama provinsi",
		},
		"city": &graphql.Field{
			Type:        graphql.String,
			Description: "nama kota",
		},
		"address": &graphql.Field{
			Type:        graphql.String,
			Description: "detail alamat",
		},
		"gender": &graphql.Field{
			Type:        graphql.String,
			Description: "jenis kelamin",
		},
		"postal_code": &graphql.Field{
			Type:        graphql.Int,
			Description: "kode pos",
		},
		"profile_progress": &graphql.Field{
			Type:        graphql.Int,
			Description: "progress profil",
		},
		"birth_at": &graphql.Field{
			Type:        graphql.String,
			Description: "tanggal lahir",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userId := p.Source.(model.User_information).User_id

				var user model.User_information
				Database.First(&user, userId)
				log.Print(user.Birth_at)

				return nil, nil
			},
		},
	},
})

var UserFileType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserFile",
	Description: "Object User file untuk menyimpan file user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID dari file",
		},
		"purpose": &graphql.Field{
			Type:        graphql.String,
			Description: "tujuan file",
		},
		"type": &graphql.Field{
			Type:        graphql.String,
			Description: "tipe file",
		},
		"path": &graphql.Field{
			Type:        graphql.String,
			Description: "file path",
		},
	},
})

var UserJobType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserJob",
	Description: "Object UserJob untuk menyimpan data pekerjaan user",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID user",
		},
		"npwp": &graphql.Field{
			Type:        graphql.String,
			Description: "nomor npwp",
		},
		"job_type": &graphql.Field{
			Type:        graphql.String,
			Description: "tipe pekerjaan",
		},
		"grade": &graphql.Field{
			Type:        graphql.String,
			Description: "level pekerjaan",
		},
		"income": &graphql.Field{
			Type:        graphql.String,
			Description: "income user",
		},
	},
})

var UserSocmedType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserSocmed",
	Description: "Object UserSocmed untuk menyimpan data social media user",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID user",
		},
		"facebook": &graphql.Field{
			Type:        graphql.String,
			Description: "akun facebook",
		},
		"instagram": &graphql.Field{
			Type:        graphql.String,
			Description: "akun instagram",
		},
		"linkedin": &graphql.Field{
			Type:        graphql.String,
			Description: "akun linkedin",
		},
	},
})

var UserBankType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserBank",
	Description: "Object UserBank untuk menyimpan data rekening bank user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID rekening bank",
		},
		"user_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "user ID",
		},
		"bank": &graphql.Field{
			Type:        graphql.String,
			Description: "nama bank",
		},
		"branch": &graphql.Field{
			Type:        graphql.String,
			Description: "cabang bank",
		},
		"account_number": &graphql.Field{
			Type:        graphql.String,
			Description: "nomor rekening bank",
		},
		"holdername": &graphql.Field{
			Type:        graphql.String,
			Description: "nama pemegang rekening",
		},
		"is_primary": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "status rekening (true jika merupakan rek. utama)",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Source.(model.User_bank).ID

				var bank model.User_bank
				Database.First(&bank, id)
				if bank.Is_primary == 0 {
					return false, nil
				}

				return true, nil
			},
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Object User untuk menyimpan data user tanijoy",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID dari User",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Nama lengkap dari User",
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "Username dari User",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "Email dari User",
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "Nomor telefon dari User",
		},
		"user_information": &graphql.Field{
			Type:        UserInformationType,
			Description: "informasi detail user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var user model.User_information
				Database.First(&user, userId)

				return user, nil
			},
		},
		"profile_picture": &graphql.Field{
			Type:        UserFileType,
			Description: "kumpulan file user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var users []model.User_file
				rows, _ := Database.Raw("select * from user_files " +
					"where user_id = $1 " +
					"and purpose = 'Profpic'", userId).Rows()
				if rows.Next() {
					var user model.User_file
					Database.ScanRows(rows, &user)
					users = append(users, user)
				}

				return users, nil
			},
		},
		"ktp_picture": &graphql.Field{
			Type:        UserFileType,
			Description: "kumpulan file user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var users []model.User_file
				rows, _ := Database.Raw("select * from user_files " +
					"where user_id = $1 " +
					"and purpose = 'Profpic'", userId).Rows()
				if rows.Next() {
					var user model.User_file
					Database.ScanRows(rows, &user)
					users = append(users, user)
				}

				return users, nil
			},
		},
		"user_files": &graphql.Field{
			Type:        graphql.NewList(UserFileType),
			Description: "kumpulan file user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var users []model.User_file
				rows, _ := Database.Raw("select * from user_files " +
					"where user_id = $1 " +
					"and purposes = 'Other'", userId).Rows()
				if rows.Next() {
					var user model.User_file
					Database.ScanRows(rows, &user)
					users = append(users, user)
				}

				return users, nil
			},
		},
		"user_bank": &graphql.Field{
			Type:        UserBankType,
			Description: "list rekening bank user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var bank model.User_bank
				Database.Where(&model.User_bank{User_id:userId}).First(&bank)

				return bank, nil
			},
		},
		"user_job": &graphql.Field{
			Type:        UserJobType,
			Description: "pekerjaan user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var job model.User_job
				Database.Where(&model.User_job{User_id:userId}).First(&job)

				return job, nil
			},
		},
		"user_socmed": &graphql.Field{
			Type:        UserSocmedType,
			Description: "Sosmed user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var sosmed model.User_socmed
				Database.Where(&model.User_socmed{User_id:userId}).First(&sosmed)

				return sosmed, nil
			},
		},
		"role": &graphql.Field{
			Type:        graphql.String,
			Description: "Nama role dari User",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get User ID
				userId := p.Source.(model.User).ID

				var roleName string
				row := Database.Raw("select name from roles where id = "+
					"(select role_id from role_user where user_id = $1)", userId).Row()
				row.Scan(&roleName)

				return roleName, nil
			},
		},
	},
})
