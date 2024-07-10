package services

import (
	"tradeMetricsCollector/internal/models"
	"tradeMetricsCollector/internal/repositories"
)

type OrderService interface {
	GetOrderBookData() ([]models.DepthOrder, error)
	SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error
	GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(r repositories.OrderRepository) OrderService {
	return &orderService{orderRepository: r}
}

func (o orderService) GetOrderBookData() ([]models.DepthOrder, error) {
	panic("implement me")
}

func (o orderService) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	panic("implement me")
}

func (o orderService) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	panic("implement me")
}

func (o orderService) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	panic("implement me")
}
