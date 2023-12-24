package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"velov-widget-api/api/handlers"
	"velov-widget-api/api/utils"
)

func main() {
	fmt.Println("🚲 Velov Widget API - Ready")

	http.HandleFunc("/station", func(w http.ResponseWriter, r *http.Request) {
		var position utils.Position

		err := json.NewDecoder(r.Body).Decode(&position)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		nearestStation := utils.GetNearestStation(position)
		station, err := handlers.FetchBikeStationData(nearestStation.Number)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(station)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
