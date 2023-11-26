package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	DBClient *DBClient
}

type ProblemDetail struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (h *Handler) getMessageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")

	messages, err := h.DBClient.GetAllMessages(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		pd := ProblemDetail{
			Type:   "about:blank",
			Title:  "Not Found",
			Status: http.StatusNotFound,
			Detail: fmt.Sprintf("No messages found for language: %s", lang),
		}
		writeProblemDetail(w, pd)
		return
	}

	writeJSON(w, messages)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func writeProblemDetail(w http.ResponseWriter, pd ProblemDetail) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(pd.Status)
	json.NewEncoder(w).Encode(pd)
}