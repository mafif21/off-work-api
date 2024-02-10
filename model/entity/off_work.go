package entity

import "time"

type OffWork struct {
	Id         int
	Name       string
	Position   string
	Start_date time.Time
	End_date   time.Time
	Status     string
}
