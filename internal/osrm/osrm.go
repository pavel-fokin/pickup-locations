package osrm

import (
	"context"

	"pavel-fokin/pickup-locations/internal/router"
)

var _ router.OSRM = (*OSRM)(nil)

type OSRM struct{}

func New() *OSRM {
	return &OSRM{}
}

func (o *OSRM) DrivingRoute_V1(ctx context.Context) error {
	return nil
}
