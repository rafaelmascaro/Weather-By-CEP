package usecase

import (
	"math"

	"github.com/rafaelmascaro/Weather-By-CEP/internal/entity"
)

type TempOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type GetTempUseCase struct {
	LocationClient entity.LocationClientInterface
	WeatherClient  entity.WeatherClientInterface
}

func NewGetTempUseCase(
	locationClient entity.LocationClientInterface,
	weatherClient entity.WeatherClientInterface,
) *GetTempUseCase {
	return &GetTempUseCase{
		LocationClient: locationClient,
		WeatherClient:  weatherClient,
	}
}

func (g *GetTempUseCase) Execute(input string) (TempOutputDTO, error) {
	cep, err := entity.NewCEP(string(input))
	if err != nil {
		return TempOutputDTO{}, err
	}

	location, err := g.LocationClient.GetLocation(cep)
	if err != nil {
		return TempOutputDTO{}, err
	}

	tempC, err := g.WeatherClient.GetWeather(location)
	if err != nil {
		return TempOutputDTO{}, err
	}

	temp := entity.NewTemperature(tempC)

	dto := TempOutputDTO{
		TempC: math.Round(temp.TempC*10) / 10,
		TempF: math.Round(temp.TempF*10) / 10,
		TempK: math.Round(temp.TempK*10) / 10,
	}

	return dto, nil
}
