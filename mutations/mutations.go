package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"apijoy/config"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// database connection
var Database *gorm.DB

// AWS session
var Session, _ = session.NewSession(&aws.Config{
	Region:      aws.String(config.S3_REGION),
	Credentials: credentials.NewStaticCredentials(config.AWS_ACCESS_KEY_ID, config.AWS_SECRET_ACCESS_KEY, config.AWS_SESSION_TOKEN),
})

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		/* POKTAN AREA */
		// Poktan Signup
		"registerPoktan": GetSignupPoktanMutation(),

		// Poktan Acc
		"accPoktan": GetAccPoktanMutation(),

		/* USER AREA */
		// User Signup
		"signupUser": GetSignupUserMutation(),

		// User ForgotPassword
		"forgotPassword": GetSendForgotPasswordEmailMutation(),

		// Reset Password
		"resetPassword": GetResetPasswordMutation(),

		// User login
		"loginUser": GetLoginUserMutation(),

		// update profile progress
		"updateProgress": GetUpdateProfileProgressMutation(),

		// User edit information
		"editUserInformation": GetEditUserInfoMutation(),

		// User edit rekening
		"editUserRekening": GetEditUserRekeningMutation(),

		// User edit job
		"editUserJob": GetEditUserJobMutation(),

		// User edit socmed
		"editUserSocmed": GetEditUserSocmedMutation(),

		// Upload profpic
		"uploadProfilePicture": UploadProfilePictureMutation(),

		// Upload KTP
		"uploadKtp": UploadKtpMutation(),

		// User logout
		"logout": GetLogoutUserMutation(),

		// User change password
		"changePassword": GetChangePasswordMutation(),

		// User hard delete
		"delete": GetDeleteUserMutation(),

		// Email Confirmation
		"emailConfirmation": GetEmailConfirmationMutation(),

		// User choosing investment concurrent
		"chooseInvestment": GetChooseInvestmentMutation(),

		// User choosing payment method
		"choosePayment": GetChoosePaymentMutation(),

		// User confirm payment
		"paymentConfirm": GetPaymentConfirmMutation(),

		// verify Investment
		"verifyTransaction": VerifyConcurrentTransaction(),

		// invest agreement
		"investAgreement": GetUserInvestAgreementMutation(),

		// invest release
		"releaseInvestment": GetInvestReleaseMutation(),

		// add review
		"addReview": GetInvestmentReviewMutation(),

		// change is_shared review
		"sharedReview": GetChangeSharedMutation(),

		// reinvest investment
		"reinvest": GetChangeReinvestMutation(),
	}
}
