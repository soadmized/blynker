package service

import (
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Service = &loggingMiddleware{}

type loggingMiddleware struct {
	next iface.Service
}

func (l loggingMiddleware) SaveData(sensor *model.Sensor) error {
	//TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetData() *model.Sensor {
	//TODO implement me
	panic("implement me")
}
