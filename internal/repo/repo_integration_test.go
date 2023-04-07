package repo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"blynker/internal/config"
	"blynker/internal/model"
)

func TestRepo(t *testing.T) {
	data := model.Sensor{
		SensorID:    "first",
		Temperature: 13,
		Light:       777,
		Movement:    1,
		UpdatedAt:   time.Now(),
	}

	conf, err := config.Read()
	require.NoError(t, err)

	repo := New(conf)

	err = repo.StoreValues(&data)
	assert.NoError(t, err)

	res := repo.GetValues()
	assert.Equal(t, &data, res)
}
