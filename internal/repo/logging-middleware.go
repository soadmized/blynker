package repo

import (
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &loggingMiddleware{}

type loggingMiddleware struct {
	next iface.Repository
}

func (l loggingMiddleware) StoreValues(data *model.Sensor) error {
	//TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetValues() *model.Sensor {
	//TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetSensorIDs() []string {
	//TODO implement me
	panic("implement me")
}
