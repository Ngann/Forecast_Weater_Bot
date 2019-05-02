package main

import "time"

//forecast url
type ForecastData struct {
	List []List `json:"list"`
	City City `json:"city"`
}

type List struct {
	ForecastMain ForecastMain `json:"main"`
	Forecast []Forecast `json:"weather"`
	Date string `json:"dt_txt"`
}

type ForecastMain struct {
	Temp float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
}

type Forecast struct {
	ForecastMain string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type City struct {
	Name string `json:"name"`
	Country string `json:"country"`
}

type WeatherSort struct {
	Temp float64
	Day time.Weekday
	Info []Forecast
	City string
	Country string
	DayOf string
}

type HighLow struct {
	High float64
	Low float64
}

//weather url
type WeatherData struct {
	Weather []Weather `json:"weather"`
	Name string `json:"name"`
	Main Main `json:"main"`
	Sys Sys `json:"sys"`
}


type Main struct {
	Temp float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
}

type Weather struct {
	Scene string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type Sys struct {
	Country string `json:"country"`
}

