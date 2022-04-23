package domain

import (
	"github.com/google/uuid"
)

type Website struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name" validate:"required,max=255"`
	Shop          string    `json:"shop"`
	URL           string    `json:"url" validate:"required,url,max=255"`
	Type          string    `json:"type"`
	Ram           int       `json:"ram"`
	InStock       bool      `json:"in_stock"`
	PriceString   string    `json:"price_string"`
	StockNumer    int       `json:"stock_number"`
	Time          string    `json:"time"`
	UpdateCounter int       `json:"update_counter"`
	UnixTime      int64     `json:"unix_time"`
}

type Websites []Website
