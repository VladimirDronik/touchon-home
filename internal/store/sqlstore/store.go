package sqlstore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"touchon_home/internal/store"
)

type Store struct {
	db        *sql.DB
	serverRep *HomeServerRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Инициализация DashboardRepository
func (s *Store) Server() store.ServerRepository {
	if s.serverRep != nil {
		return s.serverRep
	}

	s.serverRep = &HomeServerRepository{
		store: s,
	}

	return s.serverRep
}
