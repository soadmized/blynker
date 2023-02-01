package api

import (
	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
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

func (a *API) registerHandlers() {
	a.HandleFunc("/", a.DisplayValues)
	a.HandleFunc("/get_data", a.Get)
	a.HandleFunc("/set_data", a.Save)
}

func (a *API) Get(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		sensor := a.service.GetData()
		makeResponse(w, sensor)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		makeResponse(w, "WRONG METHOD, USE POST")
	}
}

func (a *API) Save(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
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
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		makeResponse(w, "WRONG METHOD, USE POST")
	}
}

func (a *API) DisplayValues(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		data := a.service.GetData()
		makeResponse(w, data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		makeResponse(w, "WRONG METHOD, USE GET")
	}
}

func makeResponse(w http.ResponseWriter, data any) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
