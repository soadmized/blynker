package repo

import (
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &loggingMiddleware{}

type loggingMiddleware struct {
	next iface.Repository
}

func (l loggingMiddleware) StoreData(data *model.Sensor) error {
	//TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetData() *model.Sensor {
	//TODO implement me
	panic("implement me")
}
