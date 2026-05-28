package main

import (
	"encoding/json"
	"net/http"

	"github.com/lucasbx33/go_tp_meteo/structs"
)

type App struct{ store *Store }

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, ErrorResponse{Error: msg, Code: code})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.store.All())
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station not found")
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	var st structs.Station
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "invalid JSON")
		return
	}
	if a.store.Has(st.Id) {
		writeError(w, http.StatusConflict, "ID_TAKEN", "station already exists")
		return
	}
	a.store.Put(st)
	writeJSON(w, http.StatusCreated, st)
}

func (a *App) updateStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var st structs.Station
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "invalid JSON")
		return
	}
	st.Id = id
	exists := a.store.Has(id)
	a.store.Put(st)
	if exists {
		writeJSON(w, http.StatusOK, st)
	} else {
		writeJSON(w, http.StatusCreated, st)
	}
}

func (a *App) deleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !a.store.Delete(id) {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *App) listObservations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station not found")
		return
	}
	writeJSON(w, http.StatusOK, st.Observations)
}
