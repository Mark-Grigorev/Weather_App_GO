package gettingweatherdata

import (
	"fmt"
	"net/http"
)

func getweather_accu(apiKey string, latitude float32, longitude float32) (WeatherAvarageData, error) {
	var weatherDataAccuu WeatherAvarageData
	//строка для получения кода города из бд accuweather
	url_city := fmt.Sprintf("http://dataservice.accuweather.com/locations/v1/cities/geoposition/search?apikey=%s&q=%f,%f", apiKey, latitude, longitude)

	var cityKey int
	//строка для получения погодных данных
	url_data := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/{%f}?apikey={%s}", cityKey, apiKey)

	reqCity, err := http.NewRequest("GET", url_city, nil)

	if err != nil {
		return weatherDataAccuu, err

	}

	reqCity.Header.Set("Accept", "application/json")

	client := http.DefaultClient
	respCity, err := client.Do(reqCity)
	if err != nil {
		return weatherDataAccuu, err
	}

	defer respCity.Body.Close()

	reqData, err := http.NewRequest("GET", url_data, nil)
	if err != nil {
		return weatherDataAccuu, err
	}

	reqData.Header.Set("Accept", "application/json")

	respData, err := client.Do(reqData)
	if err != nil {
		return weatherDataAccuu, err
	}
	defer respData.Body.Close()

	//Парсинг данных с json

	return weatherDataAccuu, nil
}
