package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JsonStations struct {
	JsonStations []JsonStation `json:"stations"`
}

type JsonStation struct {
	Id           string             `json:"id"`
	Country      string             `json:"country"`
	Name         string             `json:"name"`
	Altitude     int                `json:"altitude_m"`
	Location     JsonLocation       `json:"location"`
	Device       JsonDevice         `json:"device"`
	Observations []JsonObservations `json:"observations"`
}

type JsonLocation struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

type JsonDevice struct {
	Model        string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}

type JsonObservations struct {
	Timestamp          string         `json:"timestamp"`
	TemperatureCelsius float32        `json:"temperature_celsius"`
	HumidityPercent    float32        `json:"humidity_percent"`
	PressureHpa        float32        `json:"pressure_hpa"`
	Wind               JsonWind       `json:"wind"`
	PrecipitationMm    float32        `json:"precipitation_mm"`
	AirQuality         JsonAirQuality `json:"air_quality"`
	Conditions         string         `json:"conditions"`
	Notes              *string        `json:"notes"`
}

type JsonWind struct {
	Speed        float64 `json:"speed_kmh"`
	DirectionDeg int     `json:"direction_deg"`
}

type JsonAirQuality struct {
	Pm25 float64 `json:"pm25"`
	Pm10 float64 `json:"pm10"`
	No2  float64 `json:"no2"`
}

var countryNameToISO = map[string]string{
	"France":    "FR",
	"Belgique":  "BE",
	"Espagne":   "ES",
	"Portugal":  "PT",
	"Italie":    "IT",
	"Allemagne": "DE",
	"Pays-Bas":  "NE",
	"Autriche":  "AU",
	"Suisse":    "CH",
	"Danemark":  "DK",
	"Suède":     "SE",
	"Norvège":   "NO",
	"Pologne":   "PL",
	"Tchéquie":  "CZ",
}

func countryToISO(name string) (string, error) {
	code, ok := countryNameToISO[name]
	if !ok {
		return "", fmt.Errorf("unknown country: %q", name)
	}
	return code, nil
}

func LoadFromJSON(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var raw JsonStations
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	stations := make([]Station, 0, len(raw.JsonStations))
	for _, js := range raw.JsonStations {
		isoCode, err := countryToISO(js.Country)
		if err != nil {
			return nil, fmt.Errorf("station %q: %w", js.Id, err)
		}

		installedOnDate, err := time.Parse("2006-01-02", js.Device.InstalledOn)
		if err != nil {
			return nil, fmt.Errorf("installed on %q: %w", js.Device.InstalledOn, err)
		}

		observations := make([]Observations, 0, len(js.Observations))
		for _, obs := range js.Observations {
			timestampDate, err := time.Parse("2006-01-02T15:04:05Z07:00", obs.Timestamp)
			if err != nil {
				return nil, fmt.Errorf("timestamp %q: %w", obs.Timestamp, err)
			}
			observations = append(observations, Observations{
				Timestamp:          timestampDate,
				TemperatureCelsius: obs.TemperatureCelsius,
				HumidityPercent:    obs.HumidityPercent,
				PressureHpa:        obs.PressureHpa,
				Wind: Wind{
					Speed:        obs.Wind.Speed,
					DirectionDeg: obs.Wind.DirectionDeg,
				},
				PrecipitationMm: obs.PrecipitationMm,
				AirQuality: AirQuality{
					Pm25: obs.AirQuality.Pm25,
					Pm10: obs.AirQuality.Pm10,
					No2:  obs.AirQuality.No2,
				},
				Conditions: obs.Conditions,
				Notes:      obs.Notes,
			})
		}

		station := Station{
			Id:       js.Id,
			Country:  isoCode,
			Name:     js.Name,
			Altitude: js.Altitude,

			Location: Location{
				Lat:  js.Location.Lat,
				Long: js.Location.Long,
			},

			Device: Device{
				Model:        js.Device.Model,
				Manufacturer: js.Device.Manufacturer,
				InstalledOn:  installedOnDate,
			},

			Observations: observations,
		}

		stations = append(stations, station)

	}

	return stations, nil
}
