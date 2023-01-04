package iface

import "blynker/internal/model"

type Repository interface {
	Save(data *model.Sensor) error
	Get() *model.Sensor
}

type Service interface {
	Set(sensor *model.Sensor) error
	Get() *model.Sensor
}
