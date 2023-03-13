package api

import (
	"context"
	"pavel-fokin/pickup-locations/internal/router"
)

// Router is an interface to finding routes to pickup locations.
type Router interface {
	FindClosest(ctx context.Context) ([]router.Route, error)
}
