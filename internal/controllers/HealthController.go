package controllers

import "net/http"

type HealthController struct {
}

func (controller *HealthController) Ping(w http.ResponseWriter, r *http.Request) {
	print("test")
	w.WriteHeader(http.StatusOK)
}
