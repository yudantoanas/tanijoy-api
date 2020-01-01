package mutations

import (
	"apijoy/config"
	"apijoy/function"
	"apijoy/model"
	"apijoy/object"
	"github.com/graphql-go/graphql"
	"strings"
	"time"
	_ "github.com/aws/aws-sdk-go/service/s3"
	"os"
	"path/filepath"
	"path"
)

func GetSignupUserMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserResult,
		Description: "Mutation untuk proses sign up user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter email user",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter password",
			},
			"phone": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter no telephone",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)
			phone := p.Args["phone"].(string)

			// hashing password
			hashPassword, _ := function.HashPassword(password)

			// random hex string
			conf_code := function.GenerateCode(10)

			name := strings.Split(email, "@")
			newUser := model.User{
				Name:              name[0],
				Email:             email,
				Password:          hashPassword,
				Confirmation_code: conf_code,
				Phone:             phone,
			}

			var result model.UserAuth
			var userEmail string
			row := Database.Raw("select email from users where email = $1 and deleted_at is null", email).Row()
			row.Scan(&userEmail)

			if len(userEmail) == 0 {
				var counter int
				raw := Database.Raw("select count(*) as deleted_user from users where email = $1 and deleted_at is not null", email).Row()
				raw.Scan(&counter)

				if counter > 0 {
					Database.Exec("update users set deleted_at = null where email = $1", email)
					result.Message = "welcome back"
					raw := Database.Raw("select id from users where email = $1 and deleted_at is null", email).Row()
					raw.Scan(&newUser.ID)
					result.User = newUser
					return result, nil
				} else {
					Database.Create(&newUser)

					// assign user with specific role
					assignRole := model.Role_user{
						User_id: newUser.ID,
						Role_id: 2,
					}

					Database.Create(&assignRole)

					// send email
					subject := "Tanijoy Account Sign Up"

					// http://tanijoy.id/verify/user?id=jhkjhjhghjg
					messageBody := config.BASE_URL + "/verify/user?id=" + newUser.Confirmation_code
					if err := function.SendEmail(subject, newUser.Email, messageBody); err != nil {
						return "Failed sending email", err
					}

					result.Message = "registration success"
					result.User = newUser
					return result, nil
				}
			} else {
				result.Message = "user already exist"
				return result, nil
			}
		},
	}
}

func GetEmailConfirmationMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserResult,
		Description: "Mutation untuk konfirmasi email setelah signup",
		Args: graphql.FieldConfigArgument{
			"confirmation_code": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter confirmation_code",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			confirmationCode := p.Args["confirmation_code"].(string)

			var user model.User
			var result model.UserAuth
			row := Database.Where(&model.User{Confirmation_code: confirmationCode, Confirmed: 0}).First(&user)
			if row.Error != nil {
				result.Message = "account already confirmed"
				return result, nil
			}

			user.Confirmed = 1
			Database.Save(&user)

			result.Message = "account confirmed"
			result.User = user

			return result, nil
		},
	}
}

func GetLoginUserMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserResult,
		Description: "Mutation untuk proses login user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "user email",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "user password",
			},
			"is_remember": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "status remember me (true ketika user memilih remember me)",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email := p.Args["email"].(string)
			password := p.Args["password"].(string)
			isRemember := p.Args["is_remember"].(bool)

			var user model.User
			var result model.UserAuth

			// checking email
			Database.Find(&user, model.User{Email: email})
			if user.Email == "" {
				result.Message = "Email doesn't exist"
				return result, nil
			}

			if user.Confirmed == 0 {
				result.Message = "User not yet confirmed. Please check your email"
				return result, nil
			}

			// checking password
			if !function.CheckPasswordHash(password, user.Password) {
				result.Message = "Invalid password"
				return result, nil
			}

			if len(user.Api_token) != 0 && len(user.Remember_token) != 0 {
				result.Message = "user already login!"
				return result, nil
			}

			user.Api_token = function.GenerateToken()
			if isRemember {
				user.Remember_token = function.GenerateToken()
			}
			Database.Save(&user)

			result.Message = "Login Success"
			result.User = user
			return result, nil
		},
	}
}

func GetEditUserInfoMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserInformationType,
		Description: "Mutation untuk edit user information",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter ID user",
			},
			"no_ktp": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nomor ktp",
			},
			"gender": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter gender",
			},
			"province": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter province",
			},
			"city": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter city",
			},
			"address": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter address",
			},
			"postal_code": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "Parameter kode pos",
			},
			"birth_at": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter birth_at",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			noKtp := p.Args["no_ktp"].(string)
			gender := p.Args["gender"].(string)
			province := p.Args["province"].(string)
			city := p.Args["city"].(string)
			address := p.Args["address"].(string)
			postal := p.Args["postal_code"].(int)
			birthAt := p.Args["birth_at"].(string)
			birth, _ := time.Parse("2006-01-02", birthAt)

			var exist bool // 1 = exists, 0 = not exists
			Database.Raw("select exists (select 1 from user_information where user_id = $1)", userId).Row().Scan(&exist)
			user := model.User_information{
				User_id:           userId,
				Identities_number: noKtp,
				Gender:            gender,
				Province:          province,
				City:              city,
				Address:           address,
				Birth_at:          birth,
				Postal_code:       postal,
			}

			if !exist {
				Database.Create(&user)
			} else {
				var userInfo model.User_information
				Database.Where(&model.User_information{User_id: userId}).First(&userInfo)
				Database.Model(&userInfo).Updates(model.User_information{
					Identities_number: noKtp,
					Gender:            gender,
					Province:          province,
					City:              city,
					Address:           address,
					Birth_at:          birth,
					Postal_code:       postal,
				})
				return userInfo, nil
			}

			return user, nil
		},
	}
}

func GetUpdateProfileProgressMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Int,
		Description: "update persentase kelengkapan profil",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "user id",
			},
			"updated_fields": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "jumlah updated field",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			update := p.Args["updated_fields"].(int)

			var userInfo model.User_information
			Database.Where(&model.User_information{User_id: userId}).First(&userInfo)

			update += userInfo.Profile_progress
			Database.Model(&userInfo).Updates(&model.User_information{Profile_progress: update})

			return update, nil
		},
	}
}

func GetEditUserRekeningMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserBankType,
		Description: "Mutation untuk edit rekening user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
			"bank": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nama bank",
			},
			"branch": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nama cabang bank",
			},
			"account_number": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nomor rekening",
			},
			"holdername": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Parameter nama pemegang rekening",
			},
			"is_primary": &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "Status rekening utama",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			bank := p.Args["bank"].(string)
			branch := p.Args["branch"].(string)
			account := p.Args["account_number"].(string)
			holder := p.Args["holdername"].(string)

			isPrimary := 0
			if p.Args["is_primary"].(bool) {
				isPrimary = 1
			}

			var exist bool // 1 = exists, 0 = not exists
			Database.Raw("select exists (select 1 from user_banks where user_id = $1)", userId).Row().Scan(&exist)
			user := model.User_bank{
				User_id:        userId,
				Bank:           bank,
				Branch:         branch,
				Account_number: account,
				Holdername:     holder,
				Is_primary:     1,
			}

			if !exist {
				Database.Create(&user)
			} else {
				var userBank model.User_bank
				Database.Where(&model.User_bank{User_id: userId}).First(&userBank)
				Database.Model(&userBank).Updates(model.User_bank{
					Bank:           bank,
					Branch:         branch,
					Account_number: account,
					Holdername:     holder,
					Is_primary:     isPrimary,
				})
				return userBank, nil
			}

			return user, nil
		},
	}
}

func GetEditUserJobMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserJobType,
		Description: "Mutation untuk edit data pekerjaan user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
			"npwp": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "nomor npwp",
			},
			"job_type": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "tipe pekerjaan",
			},
			"grade": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "level pekerjaan",
			},
			"income": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "income user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			npwp := p.Args["npwp"].(string)
			jobType := p.Args["job_type"].(string)
			grade := p.Args["grade"].(string)
			income := p.Args["income"].(string)

			var exist bool // 1 = exists, 0 = not exists
			Database.Raw("select exists (select 1 from user_job where user_id = $1)", userId).Row().Scan(&exist)
			user := model.User_job{
				User_id:  userId,
				Npwp:     npwp,
				Job_type: jobType,
				Grade:    grade,
				Income:   income,
			}

			if !exist {
				Database.Create(&user)
			} else {
				var userJob model.User_job
				Database.Where(&model.User_job{User_id: userId}).First(&userJob)
				Database.Model(&userJob).Updates(model.User_job{
					Npwp:     npwp,
					Job_type: jobType,
					Grade:    grade,
					Income:   income,
				})
				return userJob, nil
			}

			return user, nil
		},
	}
}

func GetEditUserSocmedMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserSocmedType,
		Description: "Mutation untuk edit data social media user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
			"facebook": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "akun facebook",
			},
			"instagram": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "akun instagram",
			},
			"linkedin": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "akun linkedin",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			facebook := p.Args["facebook"].(string)
			instagram := p.Args["instagram"].(string)
			linkedIn := p.Args["linkedin"].(string)

			var exist bool // 1 = exists, 0 = not exists
			Database.Raw("select exists (select 1 from user_socmed where user_id = $1)", userId).Row().Scan(&exist)
			user := model.User_socmed{
				User_id:   userId,
				Facebook:  facebook,
				Instagram: instagram,
				Linkedin:  linkedIn,
			}

			if !exist {
				Database.Create(&user)
			} else {
				var userSocmed model.User_socmed
				Database.Where(&model.User_socmed{User_id: userId}).First(&userSocmed)
				Database.Model(&userSocmed).Updates(model.User_socmed{
					Facebook:  facebook,
					Instagram: instagram,
					Linkedin:  linkedIn,
				})
				return userSocmed, nil
			}

			return user, nil
		},
	}
}

func GetChangePasswordMutation() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserResult,
		Description: "Mutation untuk edit data social media user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
			"old_password": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "password lama",
			},
			"new_password": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "password baru",
			},
			"confirm_password": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "konfirmasi password baru",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)
			old := p.Args["old_password"].(string)
			new := p.Args["new_password"].(string)
			confirm := p.Args["confirm_password"].(string)

			var user model.User
			var result model.UserAuth
			Database.First(&user, userId)

			// checking old password
			if !function.CheckPasswordHash(old, user.Password) {
				result.Message = "Invalid old password"
				return result, nil
			}

			// checking password match
			if confirm == new {
				result.Message = "Password not match"
				return result, nil
			}

			// hashing password
			hashPassword, _ := function.HashPassword(new)
			user.Password = hashPassword
			Database.Save(&user)

			result.Message = "Password changed"
			result.User = user

			return result, nil
		},
	}
}

func GetDeleteUserMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk proses delete user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)

			// generate time
			current_time := time.Now()

			Database.Exec("update users set deleted_at = $2 where id = $1", userId, current_time)
			Database.Exec("update role_user set deleted_at = $2 where user_id = $1", userId, current_time)

			return "Successfully delete user", nil
		},
	}
}

func GetLogoutUserMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk proses logout user",
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Parameter ID user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			userId := p.Args["user_id"].(int)

			var user model.User
			Database.First(&user, userId)
			user.Api_token = ""
			Database.Save(&user)

			return "logout success", nil
		},
	}
}

func GetSendForgotPasswordEmailMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk mengirim email forgot password",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter email user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email := p.Args["email"].(string)

			var user model.User
			Database.Find(&user, model.User{Email: email})
			if user.Email == "" {
				return "Email doesn't exist", nil
			}

			subject := "Forgot your password ?"
			message := config.BASE_URL + "/reset/user?id=" + string(user.Confirmation_code)

			if err := function.SendEmail(subject, email, message); err != nil {
				return "Cannot send email", err
			}

			return "Forgot password email sent", nil
		},
	}
}

func GetResetPasswordMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Mutation untuk reset password",
		Args: graphql.FieldConfigArgument{
			"confirmation_code": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter confirmation_code",
			},
			"new_password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Parameter new_password",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			code := p.Args["confirmation_code"].(string)
			newPassword := p.Args["new_password"].(string)

			hash, _ := function.HashPassword(newPassword)

			var user model.User
			if err := Database.Find(&user, model.User{Confirmation_code: code}); err.Error != nil {
				return "User not exist", err.Error
			}

			user.Password = hash
			if err := Database.Save(&user); err.Error != nil {
				return "Failed change password", err.Error
			}

			return "Reset password success", nil
		},
	}
}

func UploadProfilePictureMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "upload profile picture",
		Args: graphql.FieldConfigArgument{
			"file_path": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "path file",
			},
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "user id",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filePath := p.Args["file_path"].(string)
			userId := p.Args["user_id"].(int)

			var awsUserDirectory = "/public/users/image/"

			// Open the file for use
			file, err := os.Open(filePath)
			if err != nil {
				return "failed to open file", err
			}
			defer file.Close()

			filename := filepath.Base(file.Name())

			// Get file size and read the file content into a buffer
			fileInfo, _ := file.Stat()
			var size int64 = fileInfo.Size()
			buffer := make([]byte, size)
			file.Read(buffer)

			err = function.AddFileToS3(Session, awsUserDirectory , filename, buffer, size)
			if err != nil {
				return "upload failed", nil
			}

			userFile := model.User_file{
				User_id: userId,
				Path: "https://s3-" + config.S3_REGION + ".amazonaws.com/" +
					config.S3_BUCKET + awsUserDirectory + filename,
				Purpose: "Profpic",
				Type:    "Image",
			}
			Database.Create(&userFile)

			return "upload success", nil
		},
	}
}

func UploadKtpMutation() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "upload ktp",
		Args: graphql.FieldConfigArgument{
			"file_path": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "path file",
			},
			"user_id": &graphql.ArgumentConfig{
				Type:        graphql.Int,
				Description: "user id",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filePath := p.Args["file_path"].(string)
			userId := p.Args["user_id"].(int)

			var awsUserDirectory = "/public/users/image/"

			// Open the file for use
			file, err := os.Open(filePath)
			if err != nil {
				return "failed to open file", err
			}
			defer file.Close()

			filename := filepath.Base(file.Name())
			ktpFilename := function.GenerateEncryptedFilename(filename) + path.Ext(filename)

			// Get file size and read the file content into a buffer
			fileInfo, _ := file.Stat()
			var size int64 = fileInfo.Size()
			buffer := make([]byte, size)
			file.Read(buffer)

			err = function.AddFileToS3(Session, awsUserDirectory, ktpFilename, buffer, size)
			if err != nil {
				return "upload failed", nil
			}

			userFile := model.User_file{
				User_id: userId,
				Path: "https://s3-" + config.S3_REGION + ".amazonaws.com/" +
					config.S3_BUCKET + awsUserDirectory + ktpFilename,
				Purpose: "Identities",
				Type:    "Image",
			}
			Database.Create(&userFile)

			return "upload success", nil
		},
	}
}
