package server

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type handlerError struct {
	statusCode int
	err        error
}

type errorHandlerFunc func(wr http.ResponseWriter, req *http.Request) *handlerError

// baseHandler telemetry for http request/response
func baseHandler(l *zap.Logger) func(errorHandlerFunc) http.HandlerFunc {
	return func(handlerFn errorHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// TODO add other telemetry

			// log reqeust
			reqFields := []zap.Field{
				zap.String("url", r.URL.String()),
				zap.String("host", r.URL.Host),
				zap.String("remote_addr", r.RemoteAddr),
			}

			// request headers
			for k, v := range r.Header {
				reqFields = append(reqFields, zap.Strings(k, v))
			}

			l.Debug("received request", reqFields...)
			start := time.Now().UTC()
			if err := handlerFn(w, r); err != nil {
				l.Error("handler failed", zap.Error(err.err), zap.Int("status_code", err.statusCode))
				http.Error(w, err.err.Error(), err.statusCode)
			}

			responseFields := []zap.Field{
				zap.Duration("duration", time.Since(start)),
			}

			l.Debug("handled request", responseFields...)
		}
	}
}

// basic handler
func basicHandler() errorHandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) *handlerError {
		headers := wr.Header()
		headers["any"] = []string{`123`}

		body := []byte(`some body value`)
		if _, err := wr.Write(body); err != nil {
			return &handlerError{
				statusCode: http.StatusInternalServerError,
				err:        err,
			}
		}

		wr.WriteHeader(http.StatusOK)

		return nil
	}
}

// ListenAndServe http server configuration and serve
func ListenAndServe(opts ...Option) error {
	cfg := &Config{
		Addr:           ":3000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Logger:         zap.NewNop(),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	s := http.Server{
		Addr:           cfg.Addr,
		Handler:        baseHandler(cfg.Logger)(basicHandler()), // handler without pattern matches everything
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	err := s.ListenAndServe()
	return err
}
