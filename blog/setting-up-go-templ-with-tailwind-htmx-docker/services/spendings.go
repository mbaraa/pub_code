package services

import (
	"spendings/db"
	"time"
)

type SpendingsService struct {
	store db.SpendingsStore
}

func NewSpendingService(store db.SpendingsStore) *SpendingsService {
	return &SpendingsService{store}
}

func (s *SpendingsService) AddItem(spending db.Spending) error {
	spending.SpentAt = time.Now()
	return s.store.Insert(spending)
}

func (s *SpendingsService) ListItems() ([]db.Spending, error) {
	return s.store.GetAll()
}

func (s *SpendingsService) UpdateItem(id string, newValue db.Spending) error {
	return s.store.Update(id, newValue)
}

func (s *SpendingsService) DeleteItem(id string) error {
	return s.store.Delete(id)
}
