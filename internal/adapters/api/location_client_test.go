package api

import (
	"strings"
	"testing"

	"github.com/rafaelmascaro/Weather-By-CEP/configs"
	"github.com/stretchr/testify/assert"
)

func TestLocationClient_GetLocation(t *testing.T) {
	config, _ := configs.LoadConfig("../../../cmd/weathersystem")
	client := NewLocationClient(config.LocationClientUrl)

	city, err := client.GetLocation("13098401")
	assert.Equal(t, "CAMPINAS", strings.ToUpper(city))
	assert.NoError(t, err)

	_, err = client.GetLocation("")
	assert.Error(t, err, "invalid zipcode")

	_, err = client.GetLocation("13098401")
	assert.Error(t, err, "not exists zipcode")
}
