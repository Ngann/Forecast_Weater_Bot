package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"time"
)

type FetchForecast struct {
}

type FetchForecastInput struct {
	City string
	ApiKey string
}

//type FetchForecastOutput struct {
//	Data ForecastData
//}
type FetchForecastOutput struct {
	Data []WeatherSort
}

func (FetchForecast) Name() string {
	return "fetch_forecast"
}

func (FetchForecast) Version() string {
	return "1.0"
}

func (s FetchForecast) Execute(in step.Context) (interface{}, error) {
	input := FetchForecastInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output, err := s.execute(input)
	return output, err
}


func (FetchForecast) execute(input FetchForecastInput) (*FetchForecastOutput, error) {
	start := time.Now()
	requestUrl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?&APPID=%s&q=%s&units=imperial", input.ApiKey, input.City)
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

	data := ForecastData{}
	err = jsoniter.Unmarshal(body, &data)
	if err !=nil {
		panic(err)
	}

	var examples []WeatherSort
	const layoutDate = "2006-01-02 15:04:05"

	for _, rows := range data.List {
		var temp = math.Round(rows.ForecastMain.Temp)
		date, err := time.Parse(layoutDate, rows.Date)
		if err != nil {
			fmt.Println(err)
		}
		weekday := date.Weekday()
		info := WeatherSort{
			Temp: temp,
			Day:  weekday,
			Info: rows.Forecast,
			City: data.City.Name,
			Country: data.City.Country,
			DayOf: weekday.String(),
		}
		examples = append(examples, info)
	}

	day := func(c1, c2 *WeatherSort) bool {
		return c1.Day < c2.Day
	}
	increasingTemps := func(c1, c2 *WeatherSort) bool {
		return c1.Temp < c2.Temp
	}

	OrderedBy(day, increasingTemps, day).Sort(examples)
	fmt.Println("By day,<temps,day:", examples)

	var n, high, low float64

	for _ , v := range examples {
		if v.Day == time.Weekday(0) {
			//fmt.Println(v.Temp, v.Day)
			if v.Temp > n {
				n = v.Temp
				high = n
			}
		}
	}
	fmt.Println("The High value is : ", high )


	for _,v:=range examples {
		if v.Day == time.Weekday(0) {
			if v.Temp > n {
			} else {
				n = v.Temp
				low = n
			}
		}
	}
	fmt.Println("The low value is : ", low)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("delta: %s/n", delta)

	return &FetchForecastOutput{
		Data: examples,
	}, nil
}

type lessFunc func(p1, p2 *WeatherSort) bool

type multiSorter struct {
	changes []WeatherSort
	less    []lessFunc
}

func (ms *multiSorter) Sort(changes []WeatherSort) {
	ms.changes = changes
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}



//var n, smallest, biggest int
//x := examples
//
//for _,v:=range x {
//if v>n {
//fmt.Println(v,">",n)
//n = v
//biggest = n
//} else {
//fmt.Println(v,"<",n)
//}
//}
//
//fmt.Println("The biggest number is ", biggest)
//for _,v:=range x {
//if v>n {
//fmt.Println(v,">",n)
//} else {
//fmt.Println(v,"<",n)
//n = v
//smallest = n
//}
//}
//fmt.Println("The smallest number is ", smallest)

