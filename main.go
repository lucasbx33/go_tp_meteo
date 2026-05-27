package main

import (
	"fmt"
	"log"
)

func main() {
	stations, err := LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(stations))
	fmt.Println(len(stations[0].Observations))
}
