package queries

import (
	"apijoy/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

// AWS session
var Session, _ = session.NewSession(&aws.Config{
	Region:      aws.String(config.S3_REGION),
	Credentials: credentials.NewStaticCredentials(config.AWS_ACCESS_KEY_ID, config.AWS_SECRET_ACCESS_KEY, config.AWS_SESSION_TOKEN),
})

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		// list all users
		"allUser": GetAllUserQueries(),

		// list all poktan
		"allPoktan": GetAllPoktanQueries(),

		// list all projects
		"allProject": GetAllProjectQueries(),

		// project details
		"project": ProjectDetail(),

		// user details
		"user": GetUserDetailQueries(),

		// All concurrents
		"allConcurrent": GetAllConcurrentsQueries(),

		// All investments
		"allInvestment": GetAllInvestmentsQueries(),

		// user investment list
		"userInvestments": GetUserInvestmentsQueries(),

		// get user concurrent investmnet
		"userConcurrent": GetUserConcurrentQueries(),

		// unconfirmed transaction
		"concurrentTransaction": UnconfirmedTransaction(),

		// get all payment methods
		"allPayMethods": GetPayMethodsQueries(),

		// get all banks
		"allBanks": GetAllBanksQueries(),

		// get bank details
		"bank": GetBankQueries(),

		// list investment activities
		"userInvestmentActivities": GetInvestmentActivitiesQueries(),

		// investment activities notif
		"investmentActivitiesNotif": GetInvestmentActivitiesNotificationQueries(),

		// total investment
		"totalInvestment": GetUserTotalInvestmentQueries(),

		// amount hasil investment
		"hasilInvestment": GetUserHasilInvestmentQueries(),

		// investasi baru selesai
		"recentInvestment": GetRecentInvestmentQueries(),

		// riwayat investasi
		"historyInvestment": GetHistoryInvestmentQueries(),

		// investasi berlangsung
		"occurringInvestment": GetOccurringInvestmentQueries(),

		// hasil investment last updated
		"investmentLastUpdated": GetUserInvestmentLastUpdatedQueries(),

		// investment percentage
		"investmentPercentage": GetInvestmentPercentageQueries(),

		// get profile percentage
		"profilePercentage": GetProfilePercentageQueries(),

		// get featured project
		"featuredProject": GetFeaturedProjectsQueries(),

		// get expired time transaction
		"expiredConcurrent": GetConcurrentTimerQueries(),

		// get total transaction
		"totalTransaction": GetTotalTransactionQueries(),

		// get cancelled transaction
		"cancelledTransaction": GetCancelledTransactionQueries(),

		// get unpaid transaction
		"unpaidTransaction": GetUnpaidTransactionQueries(),

		// get unagreed transaction
		"unagreedTransaction": GetUnagreedTransactionQueries(),

		// get expired transaction
		"expiredTransaction": GetExpiredHistoryQueries(),

		// get expired transaction
		"successTransaction": GetSuccessHistoryQueries(),

		// get finished investment
		"finishedInvestment": GetFinishedInvestmentQueries(),

		// get detail updates
		"detailInvestUpdates": GetDetailInvestmentUpdateQueries(),

		// get detail investment gallery
		"detailInvestPhotos": GetDetailInvestmentPhotosQueries(),

		/*CS Section*/

		// get transaksi yang perlu diverifikasi
		"unverifiedTransaction": UnverifiedTransaction(),

		// kelompok tani yang belum diverifikasi
		"unverifiedPoktan": ListUnverifiedPoktan(),

		// tugas lain
		"otherTask": OtherTasks(),
	}
}
