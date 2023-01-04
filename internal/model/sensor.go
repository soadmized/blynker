package model

import "time"

type Sensor struct {
	Temperature float64   `json:"temperature"`
	Light       int16     `json:"light"`
	Movement    bool      `json:"movement"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
