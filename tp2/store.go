package main

import "github.com/lucasbx33/go_tp_meteo/structs"

type Store struct {
	stations map[string]structs.Station
}

func NewStore() *Store {
	return &Store{stations: make(map[string]structs.Station)}
}

func (s *Store) Put(st structs.Station) {
	s.stations[st.Id] = st
}

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (structs.Station, bool) {
	st, ok := s.stations[id]
	return st, ok
}

func (s *Store) Delete(id string) bool {
	if !s.Has(id) {
		return false
	}
	delete(s.stations, id)
	return true
}

func (s *Store) All() []structs.Station {
	result := make([]structs.Station, 0, len(s.stations))
	for _, st := range s.stations {
		result = append(result, st)
	}
	return result
}
