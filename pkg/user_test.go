package pkg

import "testing"

func TestAccess_GetUser(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	user, errRes := a.GetUser(1237812378)
	if errRes.IsError {
		t.Fatalf("error in fetching the user \n%s", errRes.Message)
	}
	if user.ID != 1237812378 {
		t.Fatalf("error in fetching the user id, expected : 1237812378, got :%d", user.ID)
	}
	t.Log("Test get user successful")
}

func TestAccess_ChangePassword(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	errRes := a.ChangePassword(1237812378, "oLdPaSs", "nEwPaSs")
	if errRes.IsError {
		t.Fatalf("password changing for the user failed")
	}
	t.Log("Test change password successful")
}

func TestAccess_CreateUser(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	user, errRes := a.CreateUser("New User", "newuser@example.com")
	if errRes.IsError {
		t.Fatalf("Error in creating the user")
	}
	if user.Name != "New User" {
		t.Fatalf("error in fetching the name of the user, expected: New User, got: %s", user.Name)
	}
	t.Log("Test create user successful")
}
