package models

type Table struct {
	Name   string  `json:"name"`
	Starts []int64 `json:"starts"`
	Ends   []int64 `json:"ends"`
}
