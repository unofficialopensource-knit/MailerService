package app

import (
	"net/http"

	middleware "github.com/wecoach-ai/Middleware/pkg"
)

type APIServer struct {
	Mode   string
	Server http.Server
}

var service Service

func NewAPIServer(conf HTTPConfig) APIServer {
	service = *NewService(conf)

	router := http.NewServeMux()
	router.HandleFunc(
		"POST /api/mailer/contact-us",
		middleware.EnforceJSON(middleware.LogRequestData(middleware.Cors(ContactUsHandler))),
	)
	router.HandleFunc(
		"POST /api/mailer/welcome",
		middleware.EnforceJSON(middleware.LogRequestData(middleware.Cors(WelcomeHandler))),
	)
	router.HandleFunc(
		"POST /api/mailer/password-reset",
		middleware.EnforceJSON(middleware.LogRequestData(middleware.Cors(PasswordResetHandler))),
	)
	router.HandleFunc(
		"POST /api/mailer/order",
		middleware.EnforceJSON(middleware.LogRequestData(middleware.Cors(OrderReceiptHandler))),
	)

	server := http.Server{
		Addr:    conf.BindAddress,
		Handler: router,
	}
	return APIServer{
		Mode:   conf.Environment,
		Server: server,
	}
}
