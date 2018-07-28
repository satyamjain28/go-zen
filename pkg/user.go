package pkg

// User structure for user
type User struct {
	ID     int    `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
	Role   string `json:"role,omitempty"`
}


// GetUser gets the user for the given user-id
func (a Access) GetUser(userID int) (*User, ErrorResponse) {
	var user User
	return &user, ErrorResponse{IsError: false}
}
