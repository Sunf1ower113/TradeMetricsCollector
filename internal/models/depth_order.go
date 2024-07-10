package models

type DepthOrder struct {
	Price   float64
	BaseQty float64
}

type OrderBook struct {
	id       int64
	exchange string
	pair     string
	asks     []DepthOrder
	bids     []DepthOrder
}
