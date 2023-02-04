package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppPort      int    `envconfig:"APP_PORT"`
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
