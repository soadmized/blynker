package repo

import (
	"encoding/csv"
	"log"
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
	conf config.Config
}

func NewCSVRepo() *CSVRepo {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	r := CSVRepo{conf: *conf}
	return &r
}

func (r *CSVRepo) SaveData(data *model.Sensor) error {
	r.Data = *data

	temp := strconv.FormatFloat(r.Data.Temperature, 'G', 2, 64)
	light := strconv.FormatInt(int64(r.Data.Light), 10)
	movement := strconv.FormatBool(r.Data.Movement)
	updAt := r.Data.UpdatedAt.Format(time.RFC3339)

	file, err := os.OpenFile(r.conf.CSVFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll([][]string{{updAt, temp, light, movement}})
	if err != nil {
		return err
	}

	return nil
}

func (r *CSVRepo) GetData() *model.Sensor {
	return &r.Data
}
