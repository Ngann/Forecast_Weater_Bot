package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(FetchWeather{})
	step.Register(FetchForecast{})
	step.Register(PostWeather{})
	step.Run()
}


