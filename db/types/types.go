package types

import "time"

type LastFetchedCounter struct {
	CounterType string    `db:"counter_type"`
	Value       float64   `db:"value"`
	InsertedAt  time.Time `db:"inserted_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type MarketHistory struct {
	ID           int      `db:"id"`
	Date         []byte   `db:"date"`
	ClosingPrice *float64 `db:"closing_price"`
	OpeningPrice *float64 `db:"opening_price"`
}
