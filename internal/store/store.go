package store

type Store interface {
	Server() ServerRepository
}
