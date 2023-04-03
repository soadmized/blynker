package iface

import "blynker/internal/model"

//go:generate mockery --name Repository --output ../repo --filename repo-mock.go
type Repository interface {
	StoreValues(sensor *model.Sensor) error
	GetValues() *model.Sensor
	GetSensorIDs() []string
}

//go:generate mockery --name Service --output ../service --filename service-mock.go
type Service interface {
	SaveValues(sensor *model.Sensor) error
	GetValues() *model.Sensor
	GetSensorIDs() []string
}
