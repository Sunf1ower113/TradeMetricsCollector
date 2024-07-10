package models

import "time"

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

type HistoryOrder struct {
	client_name           string
	exchange_name         string
	label                 string
	pair                  string
	side                  string
	order_type            string
	base_qty              float64
	price                 float64
	algorithm_name_placed string
	lowest_sell_prc       float64
	highest_buy_prc       float64
	commission_quote_qty  float64
	time_placed           time.Time
}