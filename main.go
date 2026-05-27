package main

import (
	"fmt"
	"log"
)

func main() {
	stations, err := LoadFromJSON("./data/weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(stations))
	fmt.Println(len(stations[0].Observations))

	stationsXml, err := LoadFromXml("./data/weather_data.xml")
	fmt.Println(len(stationsXml))
	fmt.Println(len(stationsXml[0].Observations))
}
