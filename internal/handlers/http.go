package handlers

import (
	"net/http"

	"github.com/zaibon/go-template/internal/health"
)

type Handlers struct {
	health *health.Health
	// ... other dependencies
}

func NewHandlers(health *health.Health) *Handlers {
	return &Handlers{health: health}
}

func (h *Handlers) Healthz(w http.ResponseWriter, r *http.Request) {
	h.health.HealthzHandler(w, r)
}

func (h *Handlers) Readyz(w http.ResponseWriter, r *http.Request) {
	h.health.ReadyzHandler(w, r)
}

func (h *Handlers) Startupz(w http.ResponseWriter, r *http.Request) {
	h.health.StartupHandler(w, r)
}

func (h *Handlers) SomeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Some data"))
}
