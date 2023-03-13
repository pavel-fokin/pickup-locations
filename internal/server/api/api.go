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

func StatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

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

		err := router.FindClosest(ctx)
		if err != nil {
			log.Error(ctx, err, "")
			httputil.AsErrorResponse(w, ErrUnknown, http.StatusInternalServerError)
			return
		}

		StatusOK(w, r)
	}
}
