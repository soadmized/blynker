package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/service"
)

type API struct {
	http.ServeMux
	service iface.Service
}

func New() *API {
	srv := service.New()
	a := API{service: &srv}
	a.registerHandlers()

	return &a
}

func (h *API) registerHandlers() {
	h.HandleFunc("/", h.DisplayValues)
	h.HandleFunc("/get_data", h.Get)
	h.HandleFunc("/set_data", h.Set)
}

func (h *API) Get(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		sensor := h.service.Get()
		makeResponse(writer, sensor)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		makeResponse(writer, "WRONG METHOD, USE POST")
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

	s.UpdatedAt = time.Now()
	err = h.service.Set(&s)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

func (h *API) DisplayValues(writer http.ResponseWriter, request *http.Request) {
	data := h.service.Get()
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func makeResponse(writer http.ResponseWriter, data any) {
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
