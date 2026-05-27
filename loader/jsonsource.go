package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lucasbx33/go_tp_meteo/structs"
)

var countryNameToISO = map[string]string{
	"France":    "FR",
	"Belgique":  "BE",
	"Espagne":   "ES",
	"Portugal":  "PT",
	"Italie":    "IT",
	"Allemagne": "DE",
	"Pays-Bas":  "NE",
	"Autriche":  "AT",
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

func LoadFromJSON(path string) ([]structs.Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var raw structs.JsonStations
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	stations := make([]structs.Station, 0, len(raw.JsonStations))
	for _, js := range raw.JsonStations {
		isoCode, err := countryToISO(js.Country)
		if err != nil {
			return nil, fmt.Errorf("station %q: %w", js.Id, err)
		}

		installedOnDate, err := time.Parse("2006-01-02", js.Device.InstalledOn)
		if err != nil {
			return nil, fmt.Errorf("installed on %q: %w", js.Device.InstalledOn, err)
		}

		observations := make([]structs.Observations, 0, len(js.Observations))
		for _, obs := range js.Observations {
			timestampDate, err := time.Parse("2006-01-02T15:04:05Z07:00", obs.Timestamp)
			if err != nil {
				return nil, fmt.Errorf("timestamp %q: %w", obs.Timestamp, err)
			}
			observations = append(observations, structs.Observations{
				Timestamp:          timestampDate,
				TemperatureCelsius: obs.TemperatureCelsius,
				HumidityPercent:    obs.HumidityPercent,
				PressureHpa:        obs.PressureHpa,
				Wind: structs.Wind{
					Speed:        obs.Wind.Speed,
					DirectionDeg: obs.Wind.DirectionDeg,
				},
				PrecipitationMm: obs.PrecipitationMm,
				AirQuality: structs.AirQuality{
					Pm25: obs.AirQuality.Pm25,
					Pm10: obs.AirQuality.Pm10,
					No2:  obs.AirQuality.No2,
				},
				Conditions: obs.Conditions,
				Notes:      obs.Notes,
			})
		}

		station := structs.Station{
			Id:      js.Id,
			Country: isoCode,
			Name:    js.Name,
			Altitude: js.Altitude,

			Location: structs.Location{
				Lat:  js.Location.Lat,
				Long: js.Location.Long,
			},

			Device: structs.Device{
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
