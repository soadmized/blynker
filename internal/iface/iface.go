package iface

import "blynker/internal/model"

//go:generate mockery --name Repository --output ../repo --filename repo-mock.go
type Repository interface {
	StoreData(data *model.Sensor) error
	GetData() *model.Sensor
}

//go:generate mockery --name Service --output ../service --filename service-mock.go
type Service interface {
	SaveData(sensor *model.Sensor) error
	GetData() *model.Sensor
}
