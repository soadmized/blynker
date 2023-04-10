package api

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"blynker/internal/config"
	"blynker/internal/service"
)

func TestAPI_GetData(t *testing.T) {
	conf, err := config.Read()
	require.NoError(t, err)

	srvMock := service.NewMock(t)
	router := gin.New()

	api := API{
		Router:  router,
		service: srvMock,
		conf:    conf,
	}

	fmt.Println(api)
}
