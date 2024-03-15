package db

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
	"slices"
	"sync"
	"time"
)

const dbFilePath = "./db.json"

// will be used with the json database implementations
var jsonMgr = &storeManager{}

func init() {
	_, err := os.Stat(dbFilePath)
	if os.IsNotExist(err) {
		_, err = os.Create(dbFilePath)
		if err != nil {
			panic(err)
		}
	}
}

type storeSchema struct {
	Spendings []Spending `json:"spendings"`
	Balance   int64      `json:"balance"`
}

type storeManager struct {
	mu sync.RWMutex
}

func (s *storeManager) get() (storeSchema, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var store storeSchema
	f, err := os.Open(dbFilePath)
	if err != nil {
		return storeSchema{}, err
	}
	err = json.NewDecoder(f).Decode(&store)
	if errors.Is(err, io.EOF) {
		return storeSchema{}, nil
	}
	if err != nil {
		return storeSchema{}, err
	}

	return store, nil
}

func (s *storeManager) set(store storeSchema) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	formattedJson, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(dbFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(formattedJson)
	if err != nil {
		return err
	}

	return nil
}

type SpendingsStoreJson struct {
	mu sync.RWMutex
}

func NewSpendingsStoreJson() SpendingsStore {
	return &SpendingsStoreJson{}
}

func (s *SpendingsStoreJson) Insert(spending Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	spending.Id = generateId()
	store.Spendings = append(store.Spendings, spending)
	// this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Balance -= spending.Price

	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

func generateId() string {
	sha256 := sha256.New()
	sha256.Write([]byte(time.Now().String()))
	return hex.EncodeToString(sha256.Sum(nil))

}

func (s *SpendingsStoreJson) GetAll() ([]Spending, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, err := jsonMgr.get()
	if err != nil {
		return nil, err
	}

	return store.Spendings, nil
}

func (s *SpendingsStoreJson) Update(id string, values Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	// find element by id, using the fancy `slices.IndexFunc`
	idx := slices.IndexFunc(store.Spendings, func(s Spending) bool {
		return s.Id == id
	})
	if idx == -1 {
		return errors.New("item was not found")
	}

	// update balance before the update
	// this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Balance += store.Spendings[idx].Price
	store.Balance -= values.Price

	values.Id = id
	// update the item's value
	store.Spendings[idx] = values
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpendingsStoreJson) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	// find element by id, using the fancy `slices.IndexFunc`
	idx := slices.IndexFunc(store.Spendings, func(s Spending) bool {
		return s.Id == id
	})
	if idx == -1 {
		return errors.New("item was not found")
	}

	// update balance before the deletion
	// this is bad practice updating the balance from the spendings store wrapper, but again this is just a proof of concept db
	store.Balance += store.Spendings[idx].Price

	// delete item and remove its entry from the slice
	store.Spendings = append(store.Spendings[:idx], store.Spendings[idx+1:]...)
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

type BalanceStoreJson struct {
	mu sync.RWMutex
}

func NewBalanceStoreJson() BalanceStore {
	return &BalanceStoreJson{}
}

func (b *BalanceStoreJson) GetBalance() int64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	store, err := jsonMgr.get()
	if err != nil {
		return 0
	}

	return store.Balance
}

func (b *BalanceStoreJson) SetBalance(newBalance int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	store.Balance = newBalance
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}
	return nil
}
