package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

type FetchWeather struct {
}

type FetchWeatherInput struct {
	City string
	ApiKey string
}

type FetchWeatherOutput struct {
	Data WeatherData
}

func (FetchWeather) Name() string {
	return "fetch_weather"
}

func (FetchWeather) Version() string {
	return "1.0"
}

func (s FetchWeather) Execute(in step.Context) (interface{}, error) {
	input := FetchWeatherInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output, err := s.execute(input)
	return output, err
}



func (FetchWeather) execute(input FetchWeatherInput) (*FetchWeatherOutput, error) {
	//requestUrl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?&APPID=f434e53981ac538056281117f3b69356&q=%s&units=imperial", input.City)
	requestUrl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?&APPID=%s&q=%s&units=imperial", input.ApiKey, input.City)
	//f434e53981ac538056281117f3b69356
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err, "Request: ")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, "Read: ")
		return nil, err
	}

	data := WeatherData{}

	err = jsoniter.Unmarshal(body, &data)
	if err !=nil {
		panic(err)
	}

	//const layoutDate = "2006-01-02 15:04:05"
	//now := time.Now()
	//fmt.Println(now)
	//
	//date := data.List[0].Date
	//fmt.Println(date)
	//fmt.Println(reflect.TypeOf(date))
	//
	//t,err := time.Parse(layoutDate, date)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)

	//start := time.Now()
	//end := time.Now()
	//delta := end.Sub(start)
	//
	//fmt.Printf("delta: %s/n", delta)

	//var city = data.City.Name
	//var temp = data.List[0].Main.Temp
	//var date = data.List[0].Date

	//fmt.Println("Temp:", data.List[0].Main.Temp)
	//fmt.Println("Weather Description:", data.List[0].Weather[0].Description)
	//fmt.Println("City:", data.City.Name)
	//fmt.Println("date:", data.List[0].Date)

	return &FetchWeatherOutput{
		Data: data,
	}, nil

	//return &FetchWeatherOutput{
	//	City: city,
	//	Temp: temp,
	//	Date: date,
	//}, nil
}



