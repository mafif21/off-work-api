package web

import "time"

type OffWorkResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
	Status     string    `json:"status"`
}
