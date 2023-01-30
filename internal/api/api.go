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
	a.HandleFunc("/set_data", a.Set)
}

func (a *API) Get(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		sensor := a.service.GetData()
		makeResponse(writer, sensor)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		makeResponse(writer, "WRONG METHOD, USE POST")
	}
}

func (a *API) Set(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		dec := json.NewDecoder(request.Body)
		s := model.Sensor{}

		err := dec.Decode(&s)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		}

		s.UpdatedAt = time.Now()
		err = a.service.SaveData(&s)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
		}
		makeResponse(writer, "DATA IS SAVED!")
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		makeResponse(writer, "WRONG METHOD, USE POST")
	}
}

func (a *API) DisplayValues(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		data := a.service.GetData()
		makeResponse(writer, data)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		makeResponse(writer, "WRONG METHOD, USE GET")
	}
}

func makeResponse(writer http.ResponseWriter, data any) {
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
