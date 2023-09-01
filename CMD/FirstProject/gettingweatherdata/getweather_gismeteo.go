package gettingweatherdata

import (
	"fmt"
	"net/http"
)

func getweather_gismeteo(apiKey string, latitude float32, longitude float32) (WeatherAvarageData, error) {

	var weatherDataGisMeteo WeatherAvarageData

	url := fmt.Sprintf("https://api.gismeteo.net/v2/weather/current/?latitude=%flongitude=%f", latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return weatherDataGisMeteo, err
	}

	req.Header.Set("X-Gismeteo-Token", apiKey)

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return weatherDataGisMeteo, err
	}

	defer resp.Body.Close()

	return weatherDataGisMeteo, nil

}
