package apiserver

import (
	"net/http"
	"strconv"
)

// getParam отдает любой параметр, который был получен из GET запроса
func (s *server) getParam(w http.ResponseWriter, r *http.Request, getParam string) int {
	param, ok := r.URL.Query()[getParam]

	if !ok || len(param[0]) < 1 {
		s.message(w, r, http.StatusBadRequest, errParamsError)
		return 0
	}

	paramId, err := strconv.Atoi(param[0])
	if err != nil {
		s.message(w, r, http.StatusBadRequest, errParamsError)
		return 0
	}

	return paramId
}

// deviceCodeParam отдает код устройства, который был получен из GET запроса
func (s *server) deviceCodeParam(w http.ResponseWriter, r *http.Request) {
	code, ok := r.URL.Query()["code"]

	if !ok || len(code[0]) < 1 {
		s.message(w, r, http.StatusBadRequest, errParamsError)
		return
	}
	s.deviceCode = code[0]
}
