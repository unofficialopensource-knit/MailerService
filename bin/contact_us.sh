#!/bin/bash

curl --location --request POST 'http://localhost:8080/mail/contact-us' \
--header 'Content-Type: application/json' \
--data-raw '{
	"Name": "TestUser",
    "Email": "test@example.com",
    "ContactNumber": "1234567890",
    "UserType": "coach",
    "Message": "Hello, World!!",
}'

