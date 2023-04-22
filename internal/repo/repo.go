//nolint:exhaustruct
package repo

import (
	"sync"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var _ iface.Repository = &Repo{}

type Repo struct {
	Data   model.Sensor
	conf   *config.Config
	client influxdb2.Client
	wg     sync.WaitGroup
}

func New(conf *config.Config) *Repo {
	client := influxdb2.NewClient(conf.MakeInfluxURL(), conf.InfluxToken)
	r := Repo{
		conf:   conf,
		client: client,
	}

	return &r
}

func (r *Repo) GetValues() *model.Sensor {
	return &r.Data
}

func (r *Repo) StoreValues(sensor *model.Sensor) error {
	r.Data = *sensor

	defer r.client.Close()

	writeAPI := r.client.WriteAPI(r.conf.InfluxOrg, r.conf.InfluxBucket)

	tempP := r.prepareMeasurementPoint("temperature")
	lightP := r.prepareMeasurementPoint("light")
	moveP := r.prepareMeasurementPoint("movement")

	r.wg.Add(1)

	go func() {
		writeAPI.WritePoint(tempP)

		defer r.wg.Done()
	}()

	r.wg.Add(1)

	go func() {
		writeAPI.WritePoint(lightP)

		defer r.wg.Done()
	}()

	r.wg.Add(1)

	go func() {
		writeAPI.WritePoint(moveP)

		defer r.wg.Done()
	}()

	r.wg.Wait()

	return nil
}

func (r *Repo) GetSensorIDs() []string {
	// TODO implement me
	panic("implement me")
}

// prepareMeasurementPoint prepares data point for InfluxDB.
// Accepts only "temperature", "light", "movement" arguments.
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
