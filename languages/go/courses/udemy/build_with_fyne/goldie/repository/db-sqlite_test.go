package repository

import (
	"fmt"
	"testing"
	"time"
)

var holdingId int64 = 0

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

	holdingId = result.ID
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

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(holdingId)
	if err != nil {
		t.Error("error retrieving holding:", err)
	}

	if h.ID != holdingId {
		t.Error("id mismatch", fmt.Sprintf("Expected %d, received %d", h.ID, holdingId))
	}

	if h.Amount != 1 || h.PurchasePrice != 1000 {
		t.Error("mismatched data retrieved:", fmt.Sprintf("Expected amount of 1, got %d. Expect purchase price of 1000, got %d.", h.Amount, h.PurchasePrice))
	}

	h, err = testRepo.GetHoldingByID(3)
	if err == nil {
		t.Error("unexpected return value: Holding with id of 3 should not exist.")
	}
}
