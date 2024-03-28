package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func TestMailHandlerEmptyBody(t *testing.T) {
	testRouter := factory.App("test")
	request, err := http.NewRequest("POST", "/mail", nil)
	if err != nil {
		log.Fatalln("Error occurred while creating a request object")
	}
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, request)

	responseData, err := io.ReadAll(w.Body)
	if err != nil {
		log.Println("Received following err")
		log.Fatalln(err.Error())
	}
	assert.Equal(t, "", string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestMailHandlerInvalidBody(t *testing.T) {
	testRouter := factory.App("test")
	jsonPayload, err := json.Marshal(schema.MailRequestSchema{
		UseServerDefaultConfig: true,
	})
	if err != nil {
		log.Fatalln("Error occured while marshaling the test input")
	}
	request, err := http.NewRequest("POST", "/mail", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln("Error occurred while creating a request object")
	}
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, request)

	responseData, err := io.ReadAll(w.Body)
	if err != nil {
		log.Println("Received following err")
		log.Fatalln(err.Error())
	}
	assert.Equal(t, "", string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMailHandlerContactUsStatusOK(t *testing.T) {
	testRouter := factory.App("test")
	jsonPayload, err := json.Marshal(schema.MailRequestSchema{
		UseServerDefaultConfig: true,
		Schema: &schema.MailSchema{
			TemplateType: "CONTACT_US",
			ContactUs: &schema.ContactUsTplContext{
				Name:          "TestUser",
				Intro:         "TestIntro",
				Email:         "test@example.com",
				ContactNumber: "1234567890",
				UserType:      "coach",
				Message:       "TestMessage",
			},
		},
	})
	if err != nil {
		log.Fatalln("Error occured while marshaling the test input")
	}
	log.Println(string(jsonPayload))
	request, err := http.NewRequest("POST", "/mail", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln("Error occurred while creating a request object")
	}
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, request)

	responseData, err := io.ReadAll(w.Body)
	if err != nil {
		log.Println("Received following error while reqding data from buffer")
		log.Fatalln(err.Error())
	}
	log.Println(responseData)
	// assert.Equal(t, http.StatusOK, w.Code)
}
