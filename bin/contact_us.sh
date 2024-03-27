#!/bin/bash

curl --location --request POST 'http://localhost:8080/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
  "templateType": "CONTACT_US",
  "templateContext": {
    "name": "Test",
    "email": "test@example.com",
    "contactNumber": "1234567890",
    "userType": "coach",
    "message": "Hello, World!!"
  }
}'
