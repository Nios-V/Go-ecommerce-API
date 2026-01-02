package response

import (
	"encoding/json"
	"net/http"
)

type StandardResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type PaginationReposponse struct {
	Data     interface{} `json:"data,omitempty"`
	Total    int64       `json:"total,omitempty"`
	Page     int         `json:"page,omitempty"`
	LastPage int         `json:"last_page,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(StandardResponse{Data: data})
}

func PaginationJSON(w http.ResponseWriter, status int, data interface{}, total int64, page, lastPage int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(PaginationReposponse{
		Data:     data,
		Total:    total,
		Page:     page,
		LastPage: lastPage,
	})
}

func Error(w http.ResponseWriter, status int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(StandardResponse{Error: errMsg})
}
