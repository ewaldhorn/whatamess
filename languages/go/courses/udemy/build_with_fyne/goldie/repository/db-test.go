package repository

import (
	"time"
)

/*
 * Mock repository for testing
 */
type TestRepository struct{}

var purchaseDate = time.Now()

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (repo *TestRepository) Migrate() error {
	return nil
}

func (repo *TestRepository) InsertHolding(holding Holding) (*Holding, error) {
	return &holding, nil
}

func (repo *TestRepository) AllHoldings() ([]Holding, error) {
	var all []Holding

	h := Holding{
		Amount:        1,
		PurchaseDate:  purchaseDate,
		PurchasePrice: 1000,
	}

	all = append(all, h)

	return all, nil
}

func (repo *TestRepository) GetHoldingByID(id int64) (*Holding, error) {
	h := Holding{
		Amount:        1,
		PurchaseDate:  purchaseDate,
		PurchasePrice: 1000,
	}

	return &h, nil
}

func (repo *TestRepository) UpdateHolding(id int64, updated Holding) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
