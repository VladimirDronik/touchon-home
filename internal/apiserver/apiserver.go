package apiserver

import (
	"net/http"
	"touchon_home/internal/configurer"
	"touchon_home/internal/store/sqlstore"
)

// Запуск сервера
func Start(store *sqlstore.Store, config *configurer.Config) error {

	s := newServer(store, config)
	return http.ListenAndServe(config.BindAddr, s)
}
