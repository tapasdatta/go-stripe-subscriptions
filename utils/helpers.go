package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success string      `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJson(w, code, map[string]string{"message": message})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(Response{Success: "success", Code: code, Data: payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
