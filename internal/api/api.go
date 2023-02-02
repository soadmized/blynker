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
	a.routeHandlers()

	return &a
}

func (a *API) routeHandlers() {
	a.HandleFunc("/", a.Route)
}

func (a *API) Get(w http.ResponseWriter, req *http.Request) {
	sensor := a.service.GetData()
	makeResponse(w, sensor)
}

func (a *API) Save(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	s := model.Sensor{}

	err := dec.Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	s.UpdatedAt = time.Now()
	err = a.service.SaveData(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
	makeResponse(w, "DATA IS SAVED!")
}

func (a *API) DisplayValues(w http.ResponseWriter, req *http.Request) {
	data := a.service.GetData()
	makeResponse(w, data)
}

func makeResponse(w http.ResponseWriter, data any) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
