package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func ContactUsHandler(w http.ResponseWriter, r *http.Request) {
	var payload ContactUsInput
	if r.Body == nil {
		slog.Error("Received empty body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = service.SendContactUsMail(payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	var payload WelcomeInput

	if r.Body == nil {
		slog.Error("Received empty body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = service.SendWelcomeMail(payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func PasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	var payload PasswordResetInput

	if r.Body == nil {
		slog.Error("Received empty body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = service.SendPasswordResetMail(payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func OrderReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var payload OrderReceiptInput
	if r.Body == nil {
		slog.Error("Received empty body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = service.SendOrderStatusMail(payload)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
