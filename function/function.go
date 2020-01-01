package function

import (
	"apijoy/config"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mailgun/mailgun-go.v1"
	"log"
	math "math/rand"
	"net/http"
)

// mailgun config
var mg = mailgun.NewMailgun(config.Domain, config.PrivateKey, config.PublicKey)

// hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(GenerateCode(50)), 14)
	return string(bytes)
}

// generate confirmation code
func GenerateCode(n int) string {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
	}

	return hex.EncodeToString(bytes)
}

func GenerateEncryptedFilename(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// generate 3 digits unique code
func GenerateUniqueCode(low, hi int) int {
	return low + math.Intn(hi-low)
}

// sending email
func SendEmail(subject, recipient_mail, msg_body string) error {
	message := config.MailMessage{
		Sender:    "donotreply@" + config.MAILGUN_DOMAIN,
		Subject:   subject,
		Body:      msg_body,
		Recipient: recipient_mail,
	}

	return config.SendMessage(mg, message)
}

// upload file (code snippet)
func AddFileToS3(s *session.Session, directory, filename string, buffer []byte, size int64) error {

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(config.S3_BUCKET),
		Key:                  aws.String(directory + filename),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
