package osrm

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"pavel-fokin/pickup-locations/internal/router"
)

const baseURL = "https://router.project-osrm.org/route/v1/driving"

type Config struct {
	Timeout int `env:"OSRM_CLIENT_TIMEOUT_SECONDS" envDefault:"30"`
}

var _ router.OSRM = (*OSRM)(nil)

type OSRM struct {
	httpClient *http.Client
}

func New(cfg Config) *OSRM {
	return &OSRM{
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
	}
}

func (o *OSRM) DrivingRoute_V1(ctx context.Context, src router.Location, dst router.Location) (router.Route, error) {
	resp, err := o.httpClient.Get(
		fmt.Sprintf("%s/%f,%f;%f,%f?overview=false", baseURL, src.Lat, src.Long, dst.Lat, dst.Long),
	)
	if err != nil {
		return router.Route{}, fmt.Errorf("DrivingRoute_V1(): %w", err)
	}

	route, err := asRoute(resp.Body)
	if err != nil {
		return router.Route{}, fmt.Errorf("DrivingRoute_V1(): %w", err)
	}

	return route, nil
}
