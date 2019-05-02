package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type PostWeather struct {
}

type PostWeatherInput struct {
}

type PostWeatherOutput struct {
}

func (PostWeather) Name() string {
	return "post_weather"
}

func (PostWeather) Version() string {
	return "1.0"
}

func (s PostWeather) Execute(ctx step.Context) (interface{}, error) {
	input := PostWeatherInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output := PostWeatherOutput{}
	return output, nil
}

func (PostWeather) execute(input PostWeatherInput) (*PostWeatherOutput, error) {
	return &PostWeatherOutput{}, nil
}