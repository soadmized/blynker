package api

import (
	"net/http/httptest"
	"testing"
	"time"

	"blynker/internal/config"
	"blynker/internal/model"
	"blynker/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAPI_GetData(t *testing.T) {
	t.Parallel()

	conf, err := config.Read()
	require.NoError(t, err)

	time := time.Now()
	data := &model.Sensor{
		SensorID:    "sensor",
		Temperature: 12,
		Light:       666,
		Movement:    1,
		UpdatedAt:   time,
	}
	srvMock := service.NewMock(t)
	srvMock.On("GetValues").Return(data)

	router := gin.New()

	api := API{
		Router:  router,
		service: srvMock,
		conf:    conf,
	}

	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), router)
	api.GetData(ctx)
}
