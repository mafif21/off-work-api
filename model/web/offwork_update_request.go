package web

type OffworkUpdateRequest struct {
	Id     int    `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}
