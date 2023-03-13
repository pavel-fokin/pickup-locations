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

func New() *Router {
	return &Router{}
}

func (r *Router) FindClosest(ctx context.Context) error {
	err := r.osrm.DrivingRoute_V1(ctx)
	if err != nil {
		return fmt.Errorf("FindCloses(): %w", err)
	}
	return nil
}
