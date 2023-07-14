package api

// Api task response dtos live here

type TaskResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
