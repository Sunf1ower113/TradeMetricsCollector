package repositories

import (
	"database/sql"
	"tradeMetricsCollector/internal/models"
)

type OrderRepository interface {
	GetOrderBookData() ([]models.DepthOrder, error)
	SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error
	GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error)
	SaveOrder(client *models.Client, order *models.HistoryOrder) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetOrderBookData() ([]models.DepthOrder, error) {
	panic("implement me")
}
func (r *orderRepository) SaveOrderBook(exchangeName, pair string, orderBook []*models.DepthOrder) error {
	panic("implement me")
}
func (r *orderRepository) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	panic("implement me")
}
func (r *orderRepository) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	panic("implement me")
}
