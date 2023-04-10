package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/service"
)

const (
	saveDataPath    = "/save_data"
	getDataPath     = "/get_data"
	checkStatusPath = "/"
)

type API struct {
	Router  *gin.Engine
	service iface.Service
	conf    *config.Config
}

func New(conf *config.Config) *API {
	srv := service.New(conf)
	router := gin.New()
	a := &API{
		service: &srv,
		conf:    conf,
		Router:  router,
	}
	a.routeHandlers()

	return a
}

func (a *API) routeHandlers() {
	a.Router.POST(saveDataPath, a.SaveData)
	a.Router.POST(getDataPath, a.GetData)
	a.Router.GET(checkStatusPath, a.CheckStatus)
}

func (a *API) GetData(ctx *gin.Context) {
	values := a.service.GetValues()
	ctx.JSON(http.StatusOK, values)
}

func (a *API) SaveData(ctx *gin.Context) {
	dec := json.NewDecoder(ctx.Request.Body)
	s := model.Sensor{}

	err := dec.Decode(&s)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	s.UpdatedAt = time.Now()

	err = a.service.SaveValues(&s)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	ctx.String(http.StatusOK, "DATA IS SAVED!")
}

func (a *API) CheckStatus(ctx *gin.Context) {
	delta := a.service.GetValues().UpdatedAt.Sub(time.Now()).Abs()

	if delta > time.Second*5 {
		ctx.String(http.StatusOK, "Sensor is offline")
		return
	}

	ctx.String(http.StatusOK, "Sensor is online")
}
