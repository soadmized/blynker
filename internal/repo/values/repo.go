package values

import (
	"context"

	"github.com/influxdata/influxdb-client-go/v2"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &Repo{}

type Repo struct {
	Data   model.Sensor
	conf   *config.Config
	influx influxdb2.Client
}

func New(conf *config.Config) *Repo {
	r := Repo{conf: conf}
	return &r
}

func (r *Repo) GetData() *model.Sensor {
	return &r.Data
}

func (r *Repo) SaveData(data *model.Sensor) error {
	r.Data = *data

	client := influxdb2.NewClient(r.conf.MakeInfluxURL(), r.conf.InfluxToken)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(r.conf.InfluxOrg, r.conf.InfluxBucket)

	tempP := influxdb2.NewPointWithMeasurement("temperature").
		AddTag("id", data.SensorID).
		AddField("temperature", data.Temperature).
		SetTime(data.UpdatedAt)
	lightP := influxdb2.NewPointWithMeasurement("light").
		AddTag("id", data.SensorID).
		AddField("light", data.Light).
		SetTime(data.UpdatedAt)
	moveP := influxdb2.NewPointWithMeasurement("movement").
		AddTag("id", data.SensorID).
		AddField("movement", data.Movement).
		SetTime(data.UpdatedAt)

	err := writeAPI.WritePoint(context.Background(), tempP)
	if err != nil {
		return err
	}
	err = writeAPI.WritePoint(context.Background(), lightP)
	if err != nil {
		return err
	}
	err = writeAPI.WritePoint(context.Background(), moveP)
	if err != nil {
		return err
	}

	return nil
}
