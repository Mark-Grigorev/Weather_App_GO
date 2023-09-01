package gettingweatherdata

import "weather_app/CMD/FirstProject/apitokens"

type WeatherAvarageData struct {
	//температура
	Temperature float32
	//Влажность
	Humidity float32
	//Примечание
	Condition string
	//ветер и его плюшки
	//скорость м/c
	WindSpeed float32
	//скорость порывов м/с
	WindGust float32
	//Направление
	WindDir string
	//Примечание к направлению
	//«nw» — северо-западное.
	// «n» — северное.
	// «ne» — северо-восточное.
	// «e» — восточное.
	// «se» — юго-восточное.
	// «s» — южное.
	// «sw» — юго-западное.
	// «w» — западное.
	// «c» — штиль.

}

func GetWeatherData(latitude float32, longitude float32) (WeatherAvarageData, error) {

	apiKeyAcuu := apitokens.AccuWeatherToken()
	apiKeyGisMet := apitokens.GisMeteoToken()
	apiKeyYandex := apitokens.YandexToken()

	var weatherDataAccuu WeatherAvarageData
	var weatherDataYandex WeatherAvarageData
	var weatherDataGisMeteo WeatherAvarageData

	//Структура для готового вывода
	var weatherData WeatherAvarageData

	weatherDataAccuu, err := getweather_accu(apiKeyAcuu, latitude, longitude)
	if err != nil {
		return weatherDataAccuu, err
	}

	weatherDataGisMeteo, err1 := getweather_gismeteo(apiKeyGisMet, latitude, longitude)
	if err1 != nil {
		return weatherDataGisMeteo, err1
	}

	weatherDataYandex, err2 := getweather_yandex(apiKeyYandex, latitude, longitude)
	if err2 != nil {
		return weatherDataYandex, err2
	}

	return weatherData, nil
}
