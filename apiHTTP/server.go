package apiHTTP

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewServer(logger *zap.SugaredLogger, addr string, jwtSecret []byte) io.Closer {
	r := chi.NewRouter()

	h := Handler{
		Log:       logger,
		JWTSecret: jwtSecret,
	}
	// TODO: move routes to separate file
	r.Route("/api", func(r chi.Router) {
		// ###########################
		// ###### Slot Machines ######
		// ###########################
		r.Route("/machines", func(r chi.Router) {
			// ###########################
			// ####### ATKINS diet #######
			// ###########################
			r.Route("/atkins-diet", func(r chi.Router) {
				r.Route("/spins", func(r chi.Router) {
					r.Post("/", h.Spins)
				})
			})
		})
	})

	s := srv{&http.Server{
		Addr:      addr,
		Handler:   r,
		TLSConfig: nil,
	}}
	s.srv.SetKeepAlivesEnabled(false)

	go func() {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(errors.Wrap(err, "failed to serve http"))
		}
	}()

	logger.Info("HTTP API started on " + addr)

	return s
}

type srv struct {
	srv *http.Server
}

// Close will try to stop server gracefully
// In any case, server will be stopped in up to 1 minute
func (s srv) Close() error {
	// Context with timeout to terminate server if shutdown operation takes longer than minute
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shutdown gracefully")
	}

	return nil
}
