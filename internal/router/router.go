package router

import (
	"context"
	"fmt"
	"sort"
)

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

	for idx, dst := range dsts {
		route, err := r.osrm.DrivingRoute_V1(ctx, src, dst)
		if err != nil {
			return nil, fmt.Errorf("FindClosest(): %w", err)
		}
		routes[idx] = route
	}

	sort.Sort(ByTimeAndDistance(routes))

	return routes, nil
}
