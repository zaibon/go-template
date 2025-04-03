package health

import (
	"context"
	"net/http"
)

type Checker interface {
	Check(ctx context.Context) error
}

type Health struct {
	checkers []Checker
}

func NewHealth(checkers ...Checker) *Health {
	return &Health{
		checkers: checkers,
	}
}

func (h *Health) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.runChecks(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("NOT OK"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Health) ReadyzHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.runChecks(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("NOT READY"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("READY"))
}

func (h *Health) StartupHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.runChecks(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("NOT AVAILABLE"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("AVAILABLE"))
}

func (h *Health) runChecks(ctx context.Context) error {
	for _, checker := range h.checkers {
		if err := checker.Check(ctx); err != nil {
			return err
		}
	}
	return nil
}
