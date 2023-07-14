package repository

import (
	"database/sql"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS holdings(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		purchase_date INTEGER NOT NULL,
		purchase_price INTEGER NOT NULL
	);`

	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertHolding(holding Holding) (*Holding, error) {
	stmt := "INSERT INTO holdings(amount, purchase_date, purchase_price) values (?,?,?)"

	res, err := repo.Conn.Exec(stmt, holding.Amount, holding.PurchaseDate.Unix(), holding.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	holding.ID = id
	return &holding, nil
}

func (repo *SQLiteRepository) AllHoldings() ([]Holding, error) {
	query := "SELECT id, amount, purchase_date, purchase_price FROM holdings ORDER BY purchase_date"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var all []Holding

	for rows.Next() {
		var h Holding
		var unixTime int64

		err := rows.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}

		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}

	return all, nil
}
