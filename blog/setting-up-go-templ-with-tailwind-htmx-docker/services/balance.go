package services

import "spendings/db"

type BalanceService struct {
	store db.BalanceStore
}

func NewBalanceService(store db.BalanceStore) *BalanceService {
	return &BalanceService{store}
}

func (b *BalanceService) GetBalance() int64 {
	return b.store.GetBalance()
}
