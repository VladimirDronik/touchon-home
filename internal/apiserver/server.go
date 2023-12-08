package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"touchon_home/internal/JWTToken"
	"touchon_home/internal/configurer"
	"touchon_home/internal/model"
	"touchon_home/internal/store"
)

const (
	ctxKeyAllow ctxKey = iota
)

var (
	errTokenNotFind = errors.New("can not find token in header")
	errParamsError  = errors.New("params message")
	messDevNotFound = errors.New("not found")
)

type ctxKey int8

type server struct {
	router     *mux.Router
	logger     *logrus.Logger
	store      store.Store
	config     *configurer.Config
	userID     int
	deviceID   int
	deviceCode string
}

func newServer(store store.Store, config *configurer.Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
		config: config,
	}

	s.configeureRouter()
	s.logger.Info("API_SERVER IS RUNNING")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configeureRouter() {

	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.autenеificateUser)
	private.HandleFunc("/servers", s.getHomeServers()).Methods("GET")

}

// Проверка ключа на валидность и срок годности
func (s *server) autenеificateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			s.message(w, r, http.StatusBadRequest, errTokenNotFind)
			return
		}

		var err error
		//Проверяем не протух ли токен и извлекаем ID юзера
		s.userID, err = JWTToken.KeysExtract(r.Header["Token"][0], s.config.TokenSecret)
		if err != nil {
			s.message(w, r, http.StatusBadRequest, err)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyAllow, true)))
	})

}

func (s *server) message(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"message": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Получение всех серверов, которые доступны пользователю
func (s *server) getHomeServers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		homeServers, err := s.store.Server().GetByUser(s.userID)

		if err != nil {
			s.message(w, r, http.StatusNotFound, err)
			return
		}

		if homeServers == nil {
			s.message(w, r, http.StatusNotFound, messDevNotFound)
			return
		}

		outputJson := model.ServersStruct{Servers: homeServers}

		s.respond(w, r, http.StatusOK, outputJson)
	}
}
