package order

import (
	"net/http"
	"tradeMetricsCollector/internal/adapters/api"
	"tradeMetricsCollector/internal/services"
)

type handler struct {
	analyticsService services.OrderService
}

func (h *handler) Register(router *http.ServeMux) {

}

func NewHandler(s services.OrderService) api.Handler {
	return &handler{analyticsService: s}
}
