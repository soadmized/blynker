package config

import (
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort      int    `envconfig:"APP_PORT"`
	User         string `envconfig:"USER"`
	Pass         string `envconfig:"PASS"`
	InfluxAddr   string `envconfig:"INFLUX_ADDR"`
	InfluxPort   int    `envconfig:"INFLUX_PORT"`
	InfluxToken  string `envconfig:"INFLUX_TOKEN"`
	InfluxBucket string `envconfig:"INFLUX_BUCKET"`
	InfluxOrg    string `envconfig:"INFLUX_ORG"`
	CSVFilename  string `envconfig:"CSV_FILENAME"`
}

func Read() (*Config, error) {
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Config) MakeInfluxURL() string {
	url := c.InfluxAddr + ":" + strconv.Itoa(c.InfluxPort)
	return url
}
