package service

import (
	"blynker/internal/iface"
	"blynker/internal/model"
	"blynker/internal/repo"
)

type Service struct {
	Repo iface.Repository
}

func New() Service {
	r := repo.Repo{}
	return Service{Repo: &r}
}

func (s *Service) Set(data *model.Sensor) error {
	err := s.Repo.Save(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get() *model.Sensor {
	data := s.Repo.Get()
	return data
}
