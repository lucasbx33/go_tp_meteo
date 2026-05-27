package structs

import (
	"encoding/xml"
)

type XmlStations struct {
	XmlName     xml.Name     `xml:"weather_dataset"`
	XmlStations []XmlStation `xml:"station"`
}

type XmlStation struct {
	Id           string            `xml:"id,attr"`
	Country      string            `xml:"country,attr"`
	Name         string            `xml:"name"`
	Location     XmlLocation       `xml:"coordinates"`
	Device       XmlDevice         `xml:"hardware"`
	Observations []XmlObservations `xml:"observations>observation"`
}

type XmlLocation struct {
	Lat      float64 `xml:"lat,attr"`
	Long     float64 `xml:"lon,attr"`
	Altitude int     `xml:"altitude_m,attr"`
}

type XmlDevice struct {
	Model        string `xml:"model,attr"`
	Manufacturer string `xml:"vendor,attr"`
	InstalledOn  string `xml:"since,attr"`
}

type XmlObservations struct {
	Timestamp  string        `xml:"at,attr"`
	Wind       XmlWind       `xml:"wind"`
	AirQuality XmlAirQuality `xml:"air_quality"`
	Conditions string        `xml:"sky,attr"`
	Notes      *string       `xml:"note"`
	Measures   []xmlMeasure  `xml:"measure"`
}

type XmlWind struct {
	Speed        float64 `xml:"speed,attr"`
	DirectionDeg int     `xml:"direction,attr"`
}

type xmlMeasure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type XmlAirQuality struct {
	Pollutant []xmlPollutant `xml:"pollutant"`
}

type xmlPollutant struct {
	Type  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}
