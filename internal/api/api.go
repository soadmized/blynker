package api

import (
	"encoding/json"
	"log"
	"net/http"

	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/service"
)

type API struct {
	Router  *http.ServeMux
	service iface.Service
}

func New() *API {
	srv := service.New()
	a := API{Router: http.NewServeMux(), service: &srv}
	a.registerHandlers()

	return &a
}

func (h *API) registerHandlers() {
	h.Router.HandleFunc("/get_data", h.Get)
	h.Router.HandleFunc("/set_data", h.Set)
}

func (h *API) Get(writer http.ResponseWriter, request *http.Request) {
	sensor := h.service.Get()
	err := json.NewEncoder(writer).Encode(sensor)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

func (h *API) Set(writer http.ResponseWriter, request *http.Request) {
	dec := json.NewDecoder(request.Body)
	s := model.Sensor{}

	err := dec.Decode(&s)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	err = h.service.Set(&s)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
