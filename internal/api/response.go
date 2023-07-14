package api

type TaskResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
