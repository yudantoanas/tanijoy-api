package config

import (
	"gopkg.in/mailgun/mailgun-go.v1"
	"fmt"
)

type MailMessage struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

//Your available domain names can be found here:
//(https://app.mailgun.com/app/domains)
var Domain string = MAILGUN_DOMAIN

// The API Keys are found in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)

// starts with "key-"
var PrivateKey string = API_SECRET_KEY

// starts with "pubkey-"
var PublicKey string = PUBLIC_KEY

/*func main() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey, publicValidationKey)

	sender := "postmaster@sandbox9ff28a3c528b4c1080c6513babdbed68.mailgun.org"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "youdant@gmail.com"

	sendMessage(mg, sender, subject, body, recipient)
}*/

func SendMessage(mg mailgun.Mailgun, msg MailMessage) error {
	message := mg.NewMessage(msg.Sender, msg.Subject, msg.Body, msg.Recipient)
	resp, id, err := mg.Send(message)

	/*if err != nil {
		log.Fatal(err)
	}*/

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return err
}
