package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"velov-widget-api/api/cache"
	"velov-widget-api/api/types"
)

func FetchBikeStationData(stationNumber int) (types.BikeStation, error) {
	var station types.BikeStation
	cachedStation, _ := cache.GetCachedBikeStationData(stationNumber)

	if cachedStation.Number != 0 {
		return cachedStation, nil
	}

	resp, err := http.Get("https://api.jcdecaux.com/vls/v3/stations/" + strconv.Itoa(stationNumber) + "?contract=Lyon&apiKey=" + os.Getenv("API_KEY"))
	if err != nil {
		return types.BikeStation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.BikeStation{}, err
	}

	err = json.Unmarshal(body, &station)
	if err != nil {
		return types.BikeStation{}, err
	}

	cache.CacheBikeStationData(stationNumber, station)

	return station, nil
}
