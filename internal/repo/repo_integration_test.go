package repo

import (
	"testing"
	"time"

	"blynker/internal/config"
	"blynker/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	t.Parallel()

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
