package router

import "context"

type OSRM interface {
	DrivingRoute_V1(ctx context.Context, src Location, dst Location) (Route, error)
}
