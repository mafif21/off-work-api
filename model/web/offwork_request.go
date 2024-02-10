package web

import "time"

type OffWorkRequest struct {
	Name       string    `json:"name" validate:"required,max=255,min=1"`
	Position   string    `json:"position" validate:"required,max=255,min=1"`
	Start_date time.Time `json:"start_date" validate:"required,datetime"`
	End_date   time.Time `json:"end_date" validate:"required,datetime"`
	Status     string    `json:"status" validate:"required"`
}
