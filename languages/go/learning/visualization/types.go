package main

import "time"

// Row in CSV
type Row struct {
	Date   time.Time
	Close  float64
	Volume int
}

// Table of data
type Table struct {
	Date   []time.Time
	Price  []float64
	Volume []int
}
