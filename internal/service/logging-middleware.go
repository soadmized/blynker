//nolint:exhaustruct
package service

import (
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Service = &loggingMiddleware{}

type loggingMiddleware struct {
	next iface.Service
}

func (l loggingMiddleware) SaveValues(sensor *model.Sensor) error {
	// TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetValues() *model.Sensor {
	// TODO implement me
	panic("implement me")
}

func (l loggingMiddleware) GetSensorIDs() []string {
	// TODO implement me
	panic("implement me")
}
