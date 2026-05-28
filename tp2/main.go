package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasbx33/go_tp_meteo/loader"
)

func main() {
	stations, err := loader.LoadFromJSON("./weather_data.json")
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore()
	for _, s := range stations {
		store.Put(s)
	}
	log.Printf("bootstrap : %d stations chargées", len(stations))

	app := &App{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /stations", app.listStations)
	mux.HandleFunc("GET /stations/{id}", app.getStation)
	mux.HandleFunc("POST /stations", app.createStation)
	mux.HandleFunc("PUT /stations/{id}", app.updateStation)
	mux.HandleFunc("DELETE /stations/{id}", app.deleteStation)
	mux.HandleFunc("GET /stations/{id}/observations", app.listObservations)
	http.ListenAndServe(":8080", mux)
}
