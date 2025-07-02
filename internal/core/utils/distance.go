package utils

import (
	"math"
	"strconv"
	"velov-widget-api/internal/adapter/cache/redis"
	"velov-widget-api/internal/core/types"
)

func getStationDistance(userPosition types.Position, stationPosition types.Position) float64 {
	var latitudeDistance = math.Pow(userPosition.Latitude-stationPosition.Latitude, 2)
	var longitudeDistance = math.Pow(userPosition.Longitude-stationPosition.Longitude, 2)

	// Here math.Sqrt() is intentionally not called because we are looking for a
	// relative distance. We do not need the actual metric distance.
	return latitudeDistance + longitudeDistance
}

// GetNearestStation calculates the nearest bike station to the given user position
// It iterates through a list of stations (/resources/stations.json) and computes the distance
// Parameters:
//   - userPosition: The user's position containing latitude and longitude as JSON
//
// Returns:
//   - types.StationDistance: Object containing distance to the station and station name
func GetNearestStation(userPosition types.Position) types.StationDistance {
	closest := types.StationDistance{Station: "0", Distance: math.Inf(1)}

	for _, element := range redis.Stations {
		distance := getStationDistance(userPosition, element.Position)

		if distance < closest.Distance {
			closest = types.StationDistance{Station: strconv.Itoa(element.Number), Distance: distance}
		}
	}

	return closest
}
