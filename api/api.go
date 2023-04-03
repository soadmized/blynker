package api

import (
	"blynker/internal/config"
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
	conf    *config.Config
}

func New(conf *config.Config) *API {
	srv := service.New(conf)
	a := API{service: &srv, conf: conf}
	a.routeHandlers()

	return &a
}

func (a *API) routeHandlers() {
	a.HandleFunc("/", a.Route)
}

func (a *API) GetData(w http.ResponseWriter, req *http.Request) {
	sensor := a.service.GetValues()
	makeResponse(w, sensor)
}

func (a *API) SaveData(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	s := model.Sensor{}

	err := dec.Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	s.UpdatedAt = time.Now()
	err = a.service.SaveValues(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	makeResponse(w, "DATA IS SAVED!")
}

func (a *API) CheckStatus(w http.ResponseWriter, req *http.Request) {
	delta := a.service.GetValues().UpdatedAt.Sub(time.Now()).Abs()
	if delta > time.Second*5 {
		makeResponse(w, "Sensor is offline")
		return
	}
	makeResponse(w, "Sensor is online")
}

func makeResponse(w http.ResponseWriter, data any) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
