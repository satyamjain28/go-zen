package types

// User structure for user
type User struct {
	ID     int    `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
	Role   string `json:"role,omitempty"`
}
