package structs

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
