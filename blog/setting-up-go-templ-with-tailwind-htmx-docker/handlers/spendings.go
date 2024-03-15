package handlers

import (
	"encoding/json"
	"net/http"
	"spendings/db"
	"spendings/services"
)

type SpendingsHandler struct {
	service services.SpendingsService
}

func NewSpendingHandler(service services.SpendingsService) *SpendingsHandler {
	return &SpendingsHandler{service}
}

func (s *SpendingsHandler) HandleAddSpendingItem(w http.ResponseWriter, r *http.Request) {
	var spending db.Spending
	err := json.NewDecoder(r.Body).Decode(&spending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.AddItem(spending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}

func (s *SpendingsHandler) HandleRemoveSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exists := r.URL.Query()["id"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.service.DeleteItem(id[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}

func (s *SpendingsHandler) HandleUpdateSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exists := r.URL.Query()["id"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSpending db.Spending
	err := json.NewDecoder(r.Body).Decode(&newSpending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.service.UpdateItem(id[0], newSpending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}
