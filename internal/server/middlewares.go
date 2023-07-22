package server

import (
	"net/http"
	"time"
)

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).String()
		s.Logger.Info().Str("uri", r.RequestURI).Str("method", r.Method).Str("duration", duration).Msg("middleware!")
	})
}
