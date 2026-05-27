package main

import "time"

type Station struct {
	Id           string
	Country      string
	Name         string
	Altitude     int
	Location     Location
	Device       Device
	Observations []Observations
}

type Location struct {
	Lat  float64
	Long float64
}

type Device struct {
	Model        string
	Manufacturer string
	InstalledOn  time.Time
}

type Observations struct {
	Timestamp          time.Time
	TemperatureCelsius float32
	HumidityPercent    float32
	PressureHpa        float32
	Wind               Wind
	PrecipitationMm    float32
	AirQuality         AirQuality
	Conditions         string
	Notes              *string
}

type Wind struct {
	Speed        float64
	DirectionDeg int
}

type AirQuality struct {
	Pm25 float64
	Pm10 float64
	No2  float64
}
