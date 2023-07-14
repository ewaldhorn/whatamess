package repository

import (
	"database/sql"
	"errors"
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

func (repo *SQLiteRepository) GetHoldingByID(id int64) (*Holding, error) {
	row := repo.Conn.QueryRow("SELECT id, amount, purchase_date, purchase_price FROM holdings WHERE id = ?", id)

	var h Holding
	var unixTime int64

	err := row.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holding) error {
	if id == 0 {
		return errors.New("invalid update id (0) specified")
	}

	statement := "UPDATE holdings SET amount = ?, purchase_date = ?, purchase_price = ? WHERE id = ?"

	res, err := repo.Conn.Exec(statement, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {

}
