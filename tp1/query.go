package main

import "github.com/lucasbx33/go_tp_meteo/structs"

func FilterByCountry(stations []structs.Station, iso string) []structs.Station {
	result := make([]structs.Station, 0)
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

func AvgTemperature(s structs.Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}
	var sum float64
	for _, obs := range s.Observations {
		sum += float64(obs.TemperatureCelsius)
	}
	return sum / float64(len(s.Observations))
}

func MaxWindGust(stations []structs.Station) (structs.Station, float64) {
	var best structs.Station
	var max float64
	for _, s := range stations {
		for _, obs := range s.Observations {
			if obs.Wind.Speed > max {
				max = obs.Wind.Speed
				best = s
			}
		}
	}
	return best, max
}

func CountByCountry(stations []structs.Station) map[string]int {
	result := make(map[string]int)
	for _, s := range stations {
		result[s.Country]++
	}
	return result
}
