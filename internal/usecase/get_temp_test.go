package usecase

import (
	"errors"
	"testing"

	"github.com/rafaelmascaro/Weather-By-CEP/internal/adapters/api"
	"github.com/rafaelmascaro/Weather-By-CEP/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LocationClientMock struct {
	mock.Mock
}

func (m *LocationClientMock) GetLocation(cep entity.CEP) (string, error) {
	args := m.Called(cep)
	if cep == "13098401" {
		return "Campinas", nil
	} else {
		return "", args.Error(1)
	}
}

type WeatherClientMock struct {
}

func (m *WeatherClientMock) GetWeather(city string) (float64, error) {
	return 28.5, nil
}

func TestGetTempUseCase(t *testing.T) {
	locationClient := &LocationClientMock{}
	locationClient.On("GetLocation", entity.CEP("13098401")).Return("Campinas", nil)
	locationClient.On("GetLocation", entity.CEP("07085311")).Return("", api.ErrNotFoundZipcode)
	locationClient.On("GetLocation", entity.CEP("13098")).Return("", errors.New("bad request"))

	weatherClient := &WeatherClientMock{}

	getTemp := NewGetTempUseCase(locationClient, weatherClient)

	expected := TempOutputDTO{
		TempC: 28.5,
		TempF: 83.3,
		TempK: 301.5,
	}

	dto, err := getTemp.Execute("13098401")
	assert.Equal(t, expected, dto)
	assert.Nil(t, err)

	_, err = getTemp.Execute("07085311")
	assert.Equal(t, api.ErrNotFoundZipcode, err)

	_, err = getTemp.Execute("13098")
	assert.Equal(t, entity.ErrInvalidZipcode, err)
}
