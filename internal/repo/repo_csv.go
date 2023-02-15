package repo

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &CSVRepo{}

type CSVRepo struct {
	Data model.Sensor
	conf *config.Config
}

func NewCSVRepo(conf *config.Config) *CSVRepo {
	r := CSVRepo{conf: conf}
	return &r
}

func (r *CSVRepo) SaveData(data *model.Sensor) error {
	r.Data = *data

	sensorID := data.SensorID
	temp := strconv.FormatFloat(data.Temperature, 'G', 2, 64)
	light := strconv.FormatInt(int64(data.Light), 10)
	movement := strconv.FormatBool(data.Movement)
	updAt := data.UpdatedAt.Format(time.RFC3339)

	file, err := os.OpenFile(r.conf.CSVFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll([][]string{{sensorID, updAt, temp, light, movement}})
	if err != nil {
		return err
	}

	return nil
}

func (r *CSVRepo) GetData() *model.Sensor {
	return &r.Data
}
