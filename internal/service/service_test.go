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
	"blynker/internal/repo"
)

func TestService_GetData(t *testing.T) {
	wantData := model.Sensor{
		SensorID:    "first",
		Temperature: 13,
		Light:       666,
		Movement:    1,
		UpdatedAt:   time.Now(),
	}

	r := repo.NewMock(t)
	r.On("GetValues", mock.Anything).Return(&wantData).Once()

	conf, err := config.Read()
	require.NoError(t, err)

	srv := New(conf)
	srv.Repo = r

	res := srv.GetValues()
	assert.Equal(t, &wantData, res)
}

func TestService_SaveData(t *testing.T) {
	t.Run("positive case", func(t *testing.T) {
		wantData := model.Sensor{
			SensorID:    "first",
			Temperature: 13,
			Light:       666,
			Movement:    0,
			UpdatedAt:   time.Now(),
		}

		r := repo.NewMock(t)
		r.On("StoreValues", mock.Anything).Return(nil).Once()

		conf, err := config.Read()
		require.NoError(t, err)

		srv := New(conf)
		srv.Repo = r

		err = srv.SaveValues(&wantData)
		assert.NoError(t, err)
	})
	t.Run("repo error", func(t *testing.T) {
		wantData := model.Sensor{
			SensorID:    "first",
			Temperature: 13,
			Light:       666,
			Movement:    0,
			UpdatedAt:   time.Now(),
		}
		r := repo.NewMock(t)
		r.On("StoreValues", mock.Anything).Return(errors.New("some error")).Once()

		conf, err := config.Read()
		require.NoError(t, err)

		srv := New(conf)
		srv.Repo = r

		err = srv.SaveValues(&wantData)
		assert.Error(t, err)
	})
}
