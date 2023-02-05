package repo

import (
	"context"
	"log"

	"github.com/influxdata/influxdb-client-go/v2"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &InfluxRepo{}

type InfluxRepo struct {
	Data   model.Sensor
	conf   config.Config
	influx influxdb2.Client
}

func NewInfluxRepo() *InfluxRepo {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	r := InfluxRepo{conf: *conf}
	return &r
}

func (r *InfluxRepo) GetData() *model.Sensor {
	return &r.Data
}

func (r *InfluxRepo) SaveData(data *model.Sensor) error {
	r.Data = *data

	client := influxdb2.NewClient(r.conf.MakeInfluxURL(), r.conf.InfluxToken)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(r.conf.InfluxOrg, r.conf.InfluxBucket)

	tempP := influxdb2.NewPointWithMeasurement("temperature").
		AddTag("id", "TODO").
		AddField("temperature", data.Temperature).
		SetTime(data.UpdatedAt)
	lightP := influxdb2.NewPointWithMeasurement("light").
		AddTag("id", "TODO").
		AddField("light", data.Light).
		SetTime(data.UpdatedAt)
	moveP := influxdb2.NewPointWithMeasurement("movement").
		AddTag("id", "TODO").
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
