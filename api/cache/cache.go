package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"velov-widget-api/api/types"
)

func CacheBikeStationData(stationNumber int, station types.BikeStation) (types.BikeStation, error) {
	stationJSON, err := json.Marshal(station)
	if err != nil {
		return types.BikeStation{}, err
	}

	err = Client.Set("station:"+strconv.Itoa(stationNumber), stationJSON, 5*time.Minute).Err()

	if err != nil {
		return types.BikeStation{}, err
	}

	return station, nil
}

func GetCachedBikeStationData(stationNumber int) (types.BikeStation, error) {
	val, err := Client.Get("station:" + strconv.Itoa(stationNumber)).Result()
	if err != nil {
		return types.BikeStation{}, err
	}

	var station types.BikeStation
	err = json.Unmarshal([]byte(val), &station)
	if err != nil {
		return types.BikeStation{}, err
	}

	return station, nil
}
