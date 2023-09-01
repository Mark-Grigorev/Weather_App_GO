package main

import (
	"encoding/json"
	"net/http"
	"weather_app/CMD/FirstProject/gettingweatherdata"

	"github.com/gin-gonic/gin"
)

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

func setupRouter() *gin.Engine {
	r := gin.Default()
	var latutude float32
	var longitude float32

	r.GET("/weather", func(ctx *gin.Context) {
		WeatherData, err := gettingweatherdata.GetWeatherData(latutude, longitude)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		jsonData, err1 := json.Marshal(WeatherData)
		if err1 != nil {
			ctx.AbortWithError(http.StatusBadRequest, err1)
			return
		}
		ctx.JSON(http.StatusOK, string(jsonData))

	})

	return r
}

func main() {

}
