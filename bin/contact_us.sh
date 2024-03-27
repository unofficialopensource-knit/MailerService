#!/bin/bash

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "useServerDefaultConfig": true,
  "schema": {
    "templateType": "CONTACT_US",
    "contactUs": {
      "name": "Test",
      "email": "test@example.com",
      "contactNumber": "1234567890",
      "userType": "coach",
      "message": "Hello, World!!"
    }
  }
}'

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "useServerDefaultConfig": true,
  "schema": {
    "templateType": "WELCOME_MAIL",
    "welcomeEmail": {
      "name": "Mayank",
      "intro": "Welcome to WeCoach! We are very excited to have you on board.",
      "instruction": "To get started with WeCoach, please click here:",
      "btnColor": "#22BC66",
      "btnText": "Confirm your account",
      "btnLink": "http://wecoach.ai/login",
      "recipient": "onlinejudge95@gmail.com",
      "outro": "Need help, or have questions? Just reply to this email, we would love to help."
    }
  }
}'

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "useServerDefaultConfig": false,
  "SMTPServerConfig": {
    "identity": "",
    "username": "",
    "password": "",
    "host": "",
    "port": ""
  },
  "schema": {
    "templateType": "CONTACT_US",
    "templateContext": {
      "name": "Test",
      "email": "test@example.com",
      "contactNumber": "1234567890",
      "userType": "coach",
      "message": "Hello, World!!"
    }
  }
}'
