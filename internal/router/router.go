package router

import (
	"context"
	"fmt"
)

type Location struct {
	Long float64
	Lat  float64
}

type Route struct {
	Dst      Location
	Duration float64
	Distance float64
}

type Router struct {
	osrm OSRM
}

func New(osrm OSRM) *Router {
	return &Router{
		osrm: osrm,
	}
}

func (r *Router) FindClosest(ctx context.Context) ([]Route, error) {
	route, err := r.osrm.DrivingRoute_V1(ctx)
	if err != nil {
		return nil, fmt.Errorf("FindClosest(): %w", err)
	}
	return []Route{route}, nil
}
