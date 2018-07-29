package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

// User structure for user
type User struct {
	ID     int    `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
	Role   string `json:"role,omitempty"`
}

// UserRes structure for user response
type UserRes struct {
	User User `json:"user,omitempty"`
}

// GetUser gets the user for the given user-id
func (a Access) GetUser(userID int) (*User, ErrorResponse) {
	var userRes UserRes
	resp, errRes := NewRequest(a, fmt.Sprintf("users/%d.json", userID), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("Error in fetching the user details \n%s", errRes.Message)
		return &userRes.User, errRes
	}
	err := json.Unmarshal(resp, &userRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response on the user \n%v", err)
		return &userRes.User, ErrorResponse{IsError: true, Message: err.Error(), Code: 400}
	}
	return &userRes.User, ErrorResponse{IsError: false}
}

// CreateUser creates the user for the given name & email id
func (a Access) CreateUser(name string, emailID string) (*User, ErrorResponse) {
	var userRes UserRes
	payload := map[string]interface{}{
		"user": map[string]string{
			"name":  name,
			"email": emailID,
		},
	}
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error in marshalling the payload")
		return &userRes.User, ErrorResponse{IsError: true, Message: err.Error(), Code: 400}
	}
	resp, errRes := NewRequest(a, fmt.Sprintf("users.json"), "POST", bytePayload, 201)
	if errRes.IsError {
		log.Printf("error in creating the new user")
		return &userRes.User, errRes
	}
	err = json.Unmarshal(resp, &userRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response on the user \n%v", err)
		return &userRes.User, ErrorResponse{IsError: true, Message: err.Error(), Code: 400}
	}
	return &userRes.User, ErrorResponse{IsError: false}
}

// SetPassword sets the password for the user
func (a Access) ChangePassword(userID int, oldPassword string, newPassword string) ErrorResponse {
	payload := map[string]string{
		"previous_password": oldPassword,
		"password":          newPassword,
	}
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error in marshalling the payload")
		return ErrorResponse{IsError: true, Message: err.Error(), Code: 400}
	}
	_, errRes := NewRequest(a, fmt.Sprintf("users/%d/password.json", userID), "PUT", bytePayload, 200)
	if errRes.IsError {
		log.Printf("Error in changing the user password")
		return errRes
	}
	return ErrorResponse{IsError: false}
}
