package web

import "time"

type RoleResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
