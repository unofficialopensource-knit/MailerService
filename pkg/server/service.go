package server

import "fmt"

func SendMail(payload MailRequestSchema) {
	fmt.Println(payload)
}
