package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"velov-widget-api/api/cache"
	"velov-widget-api/api/types"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func FetchBikeStationData(ctx context.Context, stationNumber int) (types.BikeStation, error) {
	var station types.BikeStation
	cachedStation, _ := cache.GetCachedBikeStationData(ctx, stationNumber)

	if cachedStation.Number != 0 {
		return cachedStation, nil
	}

	query := url.Values{}
	query.Set("contract", "Lyon")
	query.Set("apiKey", os.Getenv("API_KEY"))
	endpoint := "https://api.jcdecaux.com/vls/v3/stations/" + strconv.Itoa(stationNumber) + "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return types.BikeStation{}, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		// Don't surface the URL (it contains the API key) to callers.
		return types.BikeStation{}, fmt.Errorf("failed to reach station provider")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.BikeStation{}, fmt.Errorf("station provider returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return types.BikeStation{}, err
	}

	err = json.Unmarshal(body, &station)
	if err != nil {
		return types.BikeStation{}, err
	}

	cache.CacheBikeStationData(ctx, stationNumber, station)

	return station, nil
}
