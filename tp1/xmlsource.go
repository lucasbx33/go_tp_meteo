package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lucasbx33/go_tp_meteo/structs"
)

func LoadFromXml(path string) ([]structs.Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var raw structs.XmlStations
	if err := xml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parsing XML: %w", err)
	}

	stations := make([]structs.Station, 0, len(raw.XmlStations))
	for _, xs := range raw.XmlStations {
		installedOnDate, err := time.Parse("2006-01-02", xs.Device.InstalledOn)
		if err != nil {
			return nil, fmt.Errorf("installed on %q: %w", xs.Device.InstalledOn, err)
		}

		observations := make([]structs.Observations, 0, len(xs.Observations))
		for _, obs := range xs.Observations {
			timestampDate, err := time.Parse("2006-01-02T15:04:05Z07:00", obs.Timestamp)
			if err != nil {
				return nil, fmt.Errorf("timestamp %q: %w", obs.Timestamp, err)
			}

			var tempC, humidity, pressure, precipitation float32
			for _, m := range obs.Measures {
				val, err := strconv.ParseFloat(m.Value, 32)
				if err != nil {
					return nil, fmt.Errorf("measure %q value %q: %w", m.Type, m.Value, err)
				}
				if m.Type == "temperature" {
					tempC = float32(val)
				} else if m.Type == "humidity" {
					humidity = float32(val)
				} else if m.Type == "pressure" {
					pressure = float32(val)
				} else if m.Type == "precipitation" {
					precipitation = float32(val)
				}
			}

			var pm25, pm10, no2 float64
			for _, p := range obs.AirQuality.Pollutant {
				val, err := strconv.ParseFloat(p.Value, 64)
				if err != nil {
					return nil, fmt.Errorf("pollutant %q value %q: %w", p.Type, p.Value, err)
				}
				if p.Type == "PM2.5" {
					pm25 = val
				} else if p.Type == "PM10" {
					pm10 = val
				} else if p.Type == "NO2" {
					no2 = val
				}
			}

			observations = append(observations, structs.Observations{
				Timestamp:          timestampDate,
				TemperatureCelsius: tempC,
				HumidityPercent:    humidity,
				PressureHpa:        pressure,
				Wind: structs.Wind{
					Speed:        obs.Wind.Speed,
					DirectionDeg: obs.Wind.DirectionDeg,
				},
				PrecipitationMm: precipitation,
				AirQuality: structs.AirQuality{
					Pm25: pm25,
					Pm10: pm10,
					No2:  no2,
				},
				Conditions: obs.Conditions,
				Notes:      obs.Notes,
			})
		}

		stations = append(stations, structs.Station{
			Id:       xs.Id,
			Country:  xs.Country,
			Name:     xs.Name,
			Altitude: xs.Location.Altitude,
			Location: structs.Location{
				Lat:  xs.Location.Lat,
				Long: xs.Location.Long,
			},
			Device: structs.Device{
				Model:        xs.Device.Model,
				Manufacturer: xs.Device.Manufacturer,
				InstalledOn:  installedOnDate,
			},
			Observations: observations,
		})
	}

	return stations, nil
}
