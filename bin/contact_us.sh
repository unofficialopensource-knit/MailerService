#!/bin/bash

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "useServerDefaultConfig": true,
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
