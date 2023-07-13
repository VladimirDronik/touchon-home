package apiserver

import (
	"net/http"
	"touchon_home/internal/configurer"
	"touchon_home/internal/store/sqlstore"
)

// Запуск сервера
func Start(store *sqlstore.Store, config *configurer.Config, certFile string, keyFile string) error {

	s := newServer(store, config)
	return http.ListenAndServeTLS(config.BindAddr, certFile, keyFile, s)
}
