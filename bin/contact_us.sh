#!/bin/bash

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "UseServerDefaultConfig": true,
  "Schema": {
    "TemplateType": "CONTACT_US",
    "ContactUs": {
      "Name": "Test",
      "Email": "test@example.com",
      "ContactNumber": "1234567890",
      "UserType": "coach",
      "Message": "Hello, World!!"
    }
  }
}'

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "UseServerDefaultConfig": true,
  "Schema": {
    "TemplateType": "WELCOME_MAIL",
    "WelcomeEmail": {
      "Name": "Mayank",
      "Intro": "Welcome to WeCoach! We are very excited to have you on board.",
      "Greeting": "Dear",
      "Instruction": "To get started with WeCoach, please click here:",
      "BtnColor": "#22BC66",
      "BtnText": "Confirm your account",
      "BtnLink": "http://wecoach.ai/login",
      "Recipient": "onlinejudge95@gmail.com",
      "Outro": "Need help, or have questions? Just reply to this email, we would love to help."
    }
  }
}'

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "UseServerDefaultConfig": false,
  "CustomSMTPConfig": {
    "Identity": "",
    "Username": "",
    "Password": "",
    "Host": "",
    "Port": ""
  },
  "Schema": {
    "TemplateType": "CONTACT_US",
    "ContactUs": {
      "Name": "Test",
      "Email": "test@example.com",
      "ContactNumber": "1234567890",
      "UserType": "coach",
      "Message": "Hello, World!!"
    }
  }
}'
