# MailerService

[![Build](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/ci.yml/badge.svg)](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/ci.yml)
[![Push](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/cd.yml/badge.svg)](https://github.com/unofficialopensource-knit/MailerService/actions/workflows/cd.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/unofficialopensource-knit/MailerService)](https://goreportcard.com/report/github.com/unofficialopensource-knit/MailerService)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/unofficialopensource-knit/MailerService)
[![codecov](https://codecov.io/gh/unofficialopensource-knit/MailerService/graph/badge.svg?token=7CIZ38MTQC)](https://codecov.io/gh/unofficialopensource-knit/MailerService)

We need to implement a backend server with the following APIs
* `POST /mail` This API will be responsible for sending out the actual email

## PreRequisite
In order to run the service locally we would be needing a set of environment variables.

* `MAILER_ENVIRONMENT`, possible values `{debug,debug-release,test,release}`.
* `MAILER_BIND_ADDR`, only to be used when running in a non-lambda based environment.
* `MAILER_SMTP_IDENTITY`, has to be an empty string `""` when using `smtp.gmail.com` for testing.
* `MAILER_SMTP_USERNAME`, username for authenticating against SMTP server.
* `MAILER_SMTP_PASSWORD`, password for authenticating against SMTP server.
* `MAILER_SMTP_HOST`, host for SMTP server.
* `MAILER_SMTP_PORT`, port for SMTP server.
* `MAILER_CONTACT_US_DEFAULT_RECIPIENT`, mail for contact us default recipients.

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
