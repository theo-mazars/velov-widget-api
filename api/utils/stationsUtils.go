package utils

import (
	"math"

	"velov-widget-api/api/cache"
)

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SelectedStation struct {
	Number   int
	Distance float64
}

func GetNearestStation(position Position) SelectedStation {
	closest := SelectedStation{Number: 0, Distance: math.Inf(1)}

	for _, element := range cache.Stations {
		distance := math.Pow(position.Latitude-element.Latitude, 2) + math.Pow(position.Longitude-element.Longitude, 2)

		if distance < closest.Distance {
			closest = SelectedStation{Number: element.Number, Distance: distance}
		}
	}

	return closest
}
