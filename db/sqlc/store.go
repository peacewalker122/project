package db

import "database/sql"

type Store struct{
	*Queries
	db *sql.DB
}

func Newstore(db *sql.DB) *Store{
	return &Store{
		Queries: New(db),
		db:      db,
	}
} 