package main

import (
	"fmt"
	"net/http"

	"github.com/rafaelmascaro/Weather-By-CEP/configs"
	"github.com/rafaelmascaro/Weather-By-CEP/internal/adapters/api"
	"github.com/rafaelmascaro/Weather-By-CEP/internal/infra/web"
	"github.com/rafaelmascaro/Weather-By-CEP/internal/infra/web/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	locationClient := api.NewLocationClient(configs.LocationClientUrl)
	weatherClient := api.NewWeatherClient(configs.WeatherClientUrl, configs.WeatherClientKey)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webTempHandler := web.NewWebTempHandler(locationClient, weatherClient)
	webserver.AddHandler("/temp", http.MethodGet, webTempHandler.Get)
	fmt.Println("Starting web server on port ", configs.WebServerPort)
	webserver.Start()
}
