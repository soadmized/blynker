package api

import (
	"blynker/internal/config"
	"blynker/internal/service"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestAPI_GetData(t *testing.T) {
	conf, err := config.Read()
	require.NoError(t, err)

	srvMock := service.NewMock(t)

	api := API{
		ServeMux: http.ServeMux{},
		service:  srvMock,
		conf:     conf,
	}

	fmt.Println(api)
}