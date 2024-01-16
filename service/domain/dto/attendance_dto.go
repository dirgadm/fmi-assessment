package dto

type AttendanceRequest struct {
	UserID    int     `json:"user_id"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type AttendanceResponse struct {
	Message string `json:"message"`
}
