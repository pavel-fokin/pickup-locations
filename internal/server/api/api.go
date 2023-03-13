package api

import (
	"errors"
	"net/http"

	"pavel-fokin/pickup-locations/internal/server/httputil"
	"pavel-fokin/pickup-locations/internal/server/log"
)

var (
	ErrValidate = errors.New("validate error")
	ErrUnknown  = errors.New(`unknown error ¯\_(ツ)_/¯`)
)

// @Summary   List of routes between source and each destination.
// @Tags      pickup-locations
// @Produce   json
// @Success   200        {object}  TODO
// @Failure   400        {object}  httputil.ErrorResponse
// @Failure   500        {object}  httputil.ErrorResponse
// @Router    /routes [get]
func RoutesGet(router Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		src, dsts, err := parseRoutesParams(r)
		if err != nil {
			log.Error(ctx, err, "Couldn't parse the route params.")
			httputil.AsErrorResponse(w, ErrValidate, http.StatusBadRequest)
			return
		}

		routes, err := router.FindClosest(ctx, src, dsts)
		if err != nil {
			log.Error(ctx, err, "Couldn't get the routes.")
			httputil.AsErrorResponse(w, ErrUnknown, http.StatusInternalServerError)
			return
		}

		resp := asRoutesGetResp(routes)
		httputil.AsSuccessResponse(w, resp, http.StatusOK)
	}
}
