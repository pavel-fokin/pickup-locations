package api

import (
	"fmt"
	"pavel-fokin/pickup-locations/internal/router"
)

type route struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type RoutesGetResp struct {
	Source string  `json:"source"`
	Routes []route `json:"routes"`
}

func asRoutesGetResp(routes []router.Route) RoutesGetResp {
	routesResp := []route{}

	for _, r := range routes {
		routesResp = append(
			routesResp,
			route{
				Destination: fmt.Sprintf("%f,%f", r.Destination.Lat, r.Destination.Long),
				Duration:    r.Duration,
				Distance:    r.Distance,
			},
		)
	}

	return RoutesGetResp{
		Routes: routesResp,
	}
}
