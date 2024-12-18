package controllers

import "net/http"

func RespondToOptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
