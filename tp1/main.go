package main

import (
	"fmt"
	"log"
)

func main() {
	stationsJSON, err := LoadFromJSON("./weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	stationsXML, err := LoadFromXml("./data/weather_data.xml")
	if err != nil {
		log.Fatal(err)
	}

	totalJSON := 0
	for _, s := range stationsJSON {
		totalJSON += len(s.Observations)
	}

	totalXML := 0
	for _, s := range stationsXML {
		totalXML += len(s.Observations)
	}

	fmt.Printf("JSON : %d stations, %d observations\n", len(stationsJSON), totalJSON)
	fmt.Printf("XML  : %d stations, %d observations\n", len(stationsXML), totalXML)

	if len(stationsJSON) == len(stationsXML) && totalJSON == totalXML {
		fmt.Println("Cohérence : OK")
	} else {
		fmt.Println("Cohérence : KO")
	}

	station, gust := MaxWindGust(stationsJSON)
	fmt.Printf("Station la plus ventée : %s (%.1f km/h)\n", station.Id, gust)

	for _, s := range stationsJSON {
		if s.Id == "FR-BOR-001" {
			fmt.Printf("Temp. moyenne Bordeaux Mérignac : %.1f °C\n", AvgTemperature(s))
			break
		}
	}

	fmt.Printf("Stations par pays : %v\n", CountByCountry(stationsJSON))
}
