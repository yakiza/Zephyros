package api

import (
	"encoding/json"
	"net/http"
)

func HealthCheckController() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(struct {
			Version   string `json:"version"`
			BuildTime string `json:"buildTime"`
			GitCommit string `json:"commit"`
		}{
			BuildTime: "",
			GitCommit: "",
			Version:   "0.0.1",
		})
	})
}
