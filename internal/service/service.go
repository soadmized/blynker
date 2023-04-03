package service

import (
	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/repo"
)

var _ iface.Service = &Service{}

type Service struct {
	Repo iface.Repository
}

func New(conf *config.Config) Service {
	r := repo.New(conf)

	return Service{Repo: r}
}

func (s *Service) SaveValues(data *model.Sensor) error {
	err := s.Repo.StoreValues(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetValues() *model.Sensor {
	data := s.Repo.GetValues()

	return data
}

func (s *Service) GetSensorIDs() []string {
	ids := s.Repo.GetSensorIDs()

	return ids
}
