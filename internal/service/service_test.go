package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"blynker/internal/config"
	"blynker/internal/model"
	"blynker/internal/repo/values"
)

func TestService_GetData(t *testing.T) {
	wantData := model.Sensor{
		SensorID:    "first",
		Temperature: 13,
		Light:       666,
		Movement:    false,
		UpdatedAt:   time.Now(),
	}

	repo := values.NewRepositoryMock(t)
	repo.On("GetData", mock.Anything).Return(&wantData).Once()

	conf, err := config.Read()
	require.NoError(t, err)

	srv := New(conf)
	srv.Repo = repo

	res := srv.GetData()
	assert.Equal(t, &wantData, res)
}

func TestService_SaveData(t *testing.T) {
	t.Run("positive case", func(t *testing.T) {
		wantData := model.Sensor{
			SensorID:    "first",
			Temperature: 13,
			Light:       666,
			Movement:    false,
			UpdatedAt:   time.Now(),
		}

		repo := values.NewRepositoryMock(t)
		repo.On("SaveData", mock.Anything).Return(nil).Once()

		conf, err := config.Read()
		require.NoError(t, err)

		srv := New(conf)
		srv.Repo = repo

		err = srv.SaveData(&wantData)
		assert.NoError(t, err)
	})
	t.Run("repo error", func(t *testing.T) {
		wantData := model.Sensor{
			SensorID:    "first",
			Temperature: 13,
			Light:       666,
			Movement:    false,
			UpdatedAt:   time.Now(),
		}
		repo := values.NewRepositoryMock(t)
		repo.On("SaveData", mock.Anything).Return(errors.New("some error")).Once()

		conf, err := config.Read()
		require.NoError(t, err)

		srv := New(conf)
		srv.Repo = repo

		err = srv.SaveData(&wantData)
		assert.Error(t, err)
	})
}
