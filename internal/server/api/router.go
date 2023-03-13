package api

import (
	"context"
)

// Router is an interface to finding routes to pickup locations.
type Router interface {
	FindClosest(ctx context.Context) error
}
