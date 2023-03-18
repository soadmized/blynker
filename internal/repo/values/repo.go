package values

import (
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/influxdata/influxdb-client-go/v2"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &Repo{}

type Repo struct {
	Data   model.Sensor
	conf   *config.Config
	client influxdb2.Client
}

func New(conf *config.Config) *Repo {
	client := influxdb2.NewClient(conf.MakeInfluxURL(), conf.InfluxToken)
	r := Repo{
		conf:   conf,
		client: client,
	}
	return &r
}

func (r *Repo) GetData() *model.Sensor {
	return &r.Data
}

func (r *Repo) SaveData(data *model.Sensor) error {
	r.Data = *data

	defer r.client.Close()

	writeAPI := r.client.WriteAPI(r.conf.InfluxOrg, r.conf.InfluxBucket)

	tempP := r.prepareMeasurementPoint("temperature")
	lightP := r.prepareMeasurementPoint("light")
	moveP := r.prepareMeasurementPoint("movement")

	go writeAPI.WritePoint(tempP)
	go writeAPI.WritePoint(lightP)
	go writeAPI.WritePoint(moveP)

	return nil
}

// prepareMeasurementPoint prepares data point for InfluxDB.
// Accepts only "temperature", "light", "movement" arguments
func (r *Repo) prepareMeasurementPoint(measurement string) *write.Point {
	var arg any

	switch measurement {
	case "temperature":
		arg = r.Data.Temperature
	case "light":
		arg = r.Data.Light
	case "movement":
		arg = r.Data.Movement
	}

	point := influxdb2.NewPointWithMeasurement(measurement).
		AddTag("id", r.Data.SensorID).
		AddField(measurement, arg).
		SetTime(r.Data.UpdatedAt)

	return point
}
