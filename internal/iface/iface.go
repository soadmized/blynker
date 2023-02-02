package iface

import "blynker/internal/model"

type Repository interface {
	SaveData(data *model.Sensor) error
	GetData() *model.Sensor
}

type Service interface {
	SaveData(sensor *model.Sensor) error
	GetData() *model.Sensor
}
