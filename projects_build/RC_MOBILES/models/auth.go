package models

type LoginRequest struct {
	Mobile string `json:"mobile" example:"9876543210"`
	Otp    string `json:"otp" example:"123456"`
}

type LoginResponse struct {
	Status string `json:"status" example:"Success"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
