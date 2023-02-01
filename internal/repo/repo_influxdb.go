package repo

import (
	"github.com/influxdata/influxdb-client-go/v2"

	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &InfluxDBRepo{}

type InfluxDBRepo struct {
	Data model.Sensor
	influxdb2.Client
}

func (r InfluxDBRepo) GetData() *model.Sensor {
	return &r.Data
}

func (r InfluxDBRepo) SaveData(data *model.Sensor) error {
	return nil
}
