package pkg

import "github.com/satyamjain28/go-zen/types"

// GetUser gets the user for the given user-id
func (a Access) GetUser(userID int) (*types.User, types.ErrorResponse) {
	var user types.User
	return &user, types.ErrorResponse{IsError: false}
}
