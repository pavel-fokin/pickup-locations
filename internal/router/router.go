package router

import (
	"context"
	"fmt"
)

type Location struct {
	Lat  float64
	Long float64
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

func (r *Router) FindClosest(ctx context.Context, src Location, dsts []Location) ([]Route, error) {
	routes := make([]Route, len(dsts))

	for _, dst := range dsts {
		route, err := r.osrm.DrivingRoute_V1(ctx, src, dst)
		if err != nil {
			return nil, fmt.Errorf("FindClosest(): %w", err)
		}
		routes = append(routes, route)
	}

	return routes, nil
}
