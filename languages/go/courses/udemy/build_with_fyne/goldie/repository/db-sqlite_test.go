package repository

import (
	"fmt"
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	purchaseDate := time.Now()
	h := Holding{
		Amount:        1,
		PurchaseDate:  purchaseDate,
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert holding failed:", err)
	}

	if result.ID <= 0 {
		t.Error("insert holding failed (no id):", err)
	}

	if result.Amount != 1 || result.PurchasePrice != 1000 {
		t.Error("insert values not as expected:", fmt.Sprintf("Amount is %d, expected %d. Price is %d, expected %d.", result.Amount, 1, result.PurchasePrice, 1000))
	}

	if result.PurchaseDate != purchaseDate {
		t.Error("insert purchase time failed:", fmt.Sprintf("got %v, expected %v", result.PurchaseDate, purchaseDate))
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("allholdings() failed:", err)
	}

	if len(h) != 1 {
		t.Error("allholdings() returned too many records:", fmt.Sprintf("Expected 1, got %d", len(h)))
	}

	if h[len(h)-1].Amount != 1 {
		t.Error("holding amount wrong:", fmt.Sprintf("Expected 1, got %d", h[len(h)-1].Amount))
	}
}
