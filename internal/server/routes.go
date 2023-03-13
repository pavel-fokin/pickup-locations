package server

import (
	"pavel-fokin/pickup-locations/internal/server/api"
)

func (s *Server) SetupRoutesAPI(router api.Router) {
	s.router.Get("/routes", api.RoutesGet(router))
}
