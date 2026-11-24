package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"velov-widget-api/api/types"
)

func CacheBikeStationData(ctx context.Context, stationNumber int, station types.BikeStation) (types.BikeStation, error) {
	stationJSON, err := json.Marshal(station)
	if err != nil {
		return types.BikeStation{}, err
	}

	err = Client.Set(ctx, "station:"+strconv.Itoa(stationNumber), stationJSON, 5*time.Minute).Err()

	if err != nil {
		return types.BikeStation{}, err
	}

	return station, nil
}

func GetCachedBikeStationData(ctx context.Context, stationNumber int) (types.BikeStation, error) {
	val, err := Client.Get(ctx, "station:"+strconv.Itoa(stationNumber)).Result()
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
