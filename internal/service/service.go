package service

import (
	"blynker/internal/config"
	"blynker/internal/iface"
	"blynker/internal/model"
	repo "blynker/internal/repo/values"
)

type Service struct {
	Repo iface.Repository
}

func New(conf *config.Config) Service {
	r := repo.New(conf)
	return Service{Repo: r}
}

func (s *Service) SaveData(data *model.Sensor) error {
	err := s.Repo.SaveData(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetData() *model.Sensor {
	data := s.Repo.GetData()
	return data
}
