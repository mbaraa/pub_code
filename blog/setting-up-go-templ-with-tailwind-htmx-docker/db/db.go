package db

import "time"

type Spending struct {
	Id      string    `json:"id"`
	Reason  string    `json:"reason"`
	Price   int64     `json:"price"`
	SpentAt time.Time `json:"spent_at"`
}

type SpendingsStore interface {
	Insert(Spending) error
	GetAll() ([]Spending, error)
	Update(id string, values Spending) error
	Delete(id string) error
}

type BalanceStore interface {
	GetBalance() int64
	SetBalance(int64) error
}
