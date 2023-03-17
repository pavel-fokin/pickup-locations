package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"pavel-fokin/pickup-locations/internal/router"
)

func parseParamToLocation(param string) (router.Location, error) {
	params := strings.Split(param, ",")

	// We expect (latitude, longtitude) pair.
	lat, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return router.Location{}, fmt.Errorf("parseParamToLocation() %w", err)
	}
	lng, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return router.Location{}, fmt.Errorf("parseParamToLocation() %w", err)
	}
	src := router.Location{
		Lat: lat,
		Lng: lng,
	}

	return src, nil
}

func parseRoutesParams(r *http.Request) (src router.Location, dsts []router.Location, err error) {
	// TODO: Add More validatons.
	srcParam := r.URL.Query().Get("src")
	if srcParam == "" {
		return router.Location{}, nil, fmt.Errorf("'src' is required")
	}

	src, err = parseParamToLocation(srcParam)
	if err != nil {
		return router.Location{}, nil, err
	}

	dstParams, ok := r.URL.Query()["dst"]
	if !ok {
		return router.Location{}, nil, fmt.Errorf("'dst' is required")
	}

	for _, param := range dstParams {
		dst, err := parseParamToLocation(param)
		if err != nil {
			return router.Location{}, nil, err
		}
		dsts = append(dsts, dst)
	}

	return src, dsts, nil
}
