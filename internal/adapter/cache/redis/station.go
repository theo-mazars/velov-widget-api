package redis

import (
	"encoding/json"
	"log/slog"
	"time"
	"velov-widget-api/internal/core/types"

	"github.com/go-redis/redis"
)

// STATION_CACHE_DURATION - The amount of time a station is kept in cache
const STATION_CACHE_DURATION = 3 * time.Minute

var Stations []types.BikeStation

// GetCacheBikeStationData - retrieves a bike station from redis cache
// Parameters:
//   - station: The key used to identify the bike station data in the cache.
//
// Returns:
//   - types.BikeStation: The bike station data.
//   - error: An error if the operation fails, otherwise nil.
func GetCacheBikeStationData(station string) (types.BikeStation, error) {
	val, err := Client.Get("station:" + station).Result()
	if err != nil {
		if err != redis.Nil { // Avoids flooding logs with record not found
			slog.Error("Unable to retrieve station from Redis", "error", err)
		}
		return types.BikeStation{}, err
	}

	var bikeStation types.BikeStation
	err = json.Unmarshal([]byte(val), &bikeStation)
	if err != nil {
		slog.Error("Unable unmarshal bikeStation JSON", "error", err)
		return types.BikeStation{}, err
	}

	return bikeStation, nil
}

// SetCacheStationData - Registers station data into Redis cache
// Parameters:
//   - station: The identifier of the station.
//   - bikeStation: JSON properties associated with the bike station.
//
// Returns:
//   - error: An error if the operation fails, otherwise nil.
func SetCacheBikeStationData(station string, bikeStation types.BikeStation) error {
	stationJSON, err := json.Marshal(bikeStation)
	if err != nil {
		return err
	}

	err = Client.Set("station:"+station, stationJSON, STATION_CACHE_DURATION).Err()
	if err != nil {
		return err
	}

	return nil
}
