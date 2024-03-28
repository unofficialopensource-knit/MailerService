# MailerService

[![Build](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/ci.yml/badge.svg)](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/ci.yml)
[![Push](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/cd.yml/badge.svg)](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/cd.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/unofficialopensource-knit/MailerService)](https://goreportcard.com/report/github.com/unofficialopensource-knit/MailerService)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/unofficialopensource-knit/MailerService)
[![codecov](https://codecov.io/gh/unofficialopensource-knit/MailerService/graph/badge.svg?token=7CIZ38MTQC)](https://codecov.io/gh/unofficialopensource-knit/MailerService)

We need to implement a backend server with the following APIs
* `GET /health` This API will be used to find the health of our server
* `POST /mail` This API will be responsible for sending out the actual email

## Development
* Use following command to install all dependencies
```bash
make install
```
* Run linting
```bash
make format
```
* Run tests
```bash
make tests
```
* Run dev server
```bash
make run-dev
```
