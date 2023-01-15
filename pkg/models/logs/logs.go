package logs

import "time"

type CreateLogsRequest struct {
	Detail    string `json:"detail"`
	CreatedBy string `json:"created_by"`
}

type CreateLogsResponse struct {
	ID        string    `json:"id"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type ListLogsRequest struct {
	Detail    string `json:"detail"`
	CreatedBy string `json:"created_by"`
}

type ListLogsResponse struct {
	ID        string `json:"id"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}
