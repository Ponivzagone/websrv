package models

import (
	"time"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type GeneralDayInfo struct {
	mainWeather string
	description string
	sunrise     time.Time
	sunset      time.Time
	icon        string
}
type WeatherTimeStamp struct {
	label    time.Time
	tempMin  float32
	tempMax  float32
	pressure float32

	windSpeed float32

	description string
	icon        string
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type OWMRequestArgs struct {
	CityId int `json:"cityId"`
}
