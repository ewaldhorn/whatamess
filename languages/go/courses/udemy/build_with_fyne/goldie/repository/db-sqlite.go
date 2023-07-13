package repository

import "database/sql"

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
