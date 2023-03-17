package osrm

import (
	"encoding/json"
	"fmt"
	"io"

	"pavel-fokin/pickup-locations/internal/router"
)

type DrivingRouteResp struct {
	Routes []struct {
		Legs []struct {
			Summary  string  `json:"summary"`
			Weight   float64 `json:"weight"`
			Duration float64 `json:"duration"`
			Steps    []any   `json:"steps"`
			Distance float64 `json:"distance"`
		} `json:"legs"`
		WeightName string  `json:"weight_name"`
		Weight     float64 `json:"weight"`
		Duration   float64 `json:"duration"`
		Distance   float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string    `json:"hint"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
	Code string `json:"code"`
}

func asRoute(body io.ReadCloser) (router.Route, error) {
	var payload DrivingRouteResp
	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return router.Route{}, fmt.Errorf("asRoute(): %w", err)
	}

	// TODO: Add validation of the OSRM response.

	// It seems that OSRM project returns (longtitude, latitude) pair.
	return router.Route{
		Destination: router.Location{
			Lng: payload.Waypoints[0].Location[0],
			Lat: payload.Waypoints[0].Location[1],
		},
		Duration: payload.Routes[0].Duration,
		Distance: payload.Routes[0].Distance,
	}, nil
}
