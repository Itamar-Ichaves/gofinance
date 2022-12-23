package db

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	dB *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		dB:      db,
		Queries: New(db),
	}
}
