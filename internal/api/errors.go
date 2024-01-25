package api

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error APIError `json:"error"`
}

func errorResponse(w http.ResponseWriter, err APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	if err.Code >= http.StatusInternalServerError {
		err.Message = "Oops, something went wrong"
	}

	errResp := ErrorResponse{Error: err}
	json.NewEncoder(w).Encode(errResp)
}
