package gettingweatherdata

import (
	"fmt"
	"net/http"
)

func getweather_yandex(apiKey string, latitude float32, longitude float32) (WeatherAvarageData, error) {

	var weatherDataYandex WeatherAvarageData
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f&lang=ru", latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return weatherDataYandex, err
	}

	req.Header.Set("X-Yandex-API-Key", apiKey)

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return weatherDataYandex, err
	}
	defer resp.Body.Close()
	return weatherDataYandex, nil
}
