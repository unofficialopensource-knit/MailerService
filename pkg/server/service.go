package server

import "fmt"

func SendMail(payload MailRequestSchema) {
	fmt.Println(payload.UseServerMail)
	fmt.Println(payload.CustomMailConfig)
	fmt.Println(payload.TemplateSchema)
	fmt.Println(payload.TemplateSchema.Context)
}
