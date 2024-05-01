package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func TestMailHandlerStatusBadRequest(t *testing.T) {
	testRouter := factory.App("test")
	assert := assert.New(t)

	t.Run("Check that test fails for empty body", func(t *testing.T) {
		request, err := http.NewRequest("POST", "/mail", nil)
		assert.Empty(err)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, request)
		response, err := io.ReadAll(w.Body)
		assert.Empty(err)

		assert.Empty(string(response))
		assert.Equal(w.Code, http.StatusBadRequest)
	})

	t.Run("Check that test fails for Schema not present", func(t *testing.T) {
		jsonPayload, err := json.Marshal(schema.MailRequestSchema{
			UseServerDefaultConfig: true,
			CustomSMTPConfig:       nil,
			Schema:                 nil,
		})
		assert.Empty(err)

		request, err := http.NewRequest("POST", "/mail", bytes.NewBuffer(jsonPayload))
		assert.Empty(err)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, request)
		response, err := io.ReadAll(w.Body)
		assert.Empty(err)

		assert.Empty(string(response))
		assert.Equal(w.Code, http.StatusBadRequest)
	})
}

// func TestMailHandlerStatusOK(t *testing.T) {
// 	testRouter := factory.App("test")
// 	assert := assert.New(t)

// 	t.Run("Check that test passes for contact us", func(t *testing.T) {
// 		jsonPayload, err := json.Marshal(schema.MailRequestSchema{
// 			UseServerDefaultConfig: true,
// 			CustomSMTPConfig:       nil,
// 			Schema: &schema.MailSchema{
// 				TemplateType: "CONTACT_US",
// 				ContactUs: &schema.ContactUsTplContext{
// 					Name:          "Test",
// 					Email:         "test@example.com",
// 					Intro:         "TestIntro",
// 					ContactNumber: "1234567890",
// 					UserType:      "coach",
// 					Message:       "Test",
// 				},
// 			},
// 		})
// 		assert.Empty(err)

// 		request, err := http.NewRequest("POST", "/mail", bytes.NewBuffer(jsonPayload))
// 		assert.Empty(err)

// 		w := httptest.NewRecorder()
// 		testRouter.ServeHTTP(w, request)
// 		response, err := io.ReadAll(w.Body)
// 		assert.Empty(err)

// 		assert.Equal(w.Code, http.StatusOK)
// 		assert.Empty(string(response))
// 	})
// }
