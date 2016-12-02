package main

import "time"

//go:generate sqlla

//+table: example
type Example struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
