package pkg

import (
	"testing"
)

func TestAccess_GetTicket(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	ticket, err := a.GetTicket(1)
	if err.IsError {
		t.Fatalf("error in fetching ticket \n%v", err)
	}
	if ticket.ID != 1 {
		t.Fatalf("unexpected ticket id, expected :%d ,got :%d", 1, ticket.ID)
	}
	t.Log("Tests for fetching ticket successful")
}

func TestAccess_GetTicketUnauthorized(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	_, err := a.GetTicket(1)
	if !err.IsError {
		t.Fatalf("error in testing authentication \n%v", err)
	}
	if err.Code != 401 {
		t.Fatalf("unexpected exception code, expected: %d, got: %d", 401, err.Code)
	}
	t.Log("Tests for authentication successful")
}

func TestAccess_ListTickets(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	_, err := a.ListTickets(1)
	if err.IsError {
		t.Fatalf("error in fetching tickets \n%v", err)
	}
	t.Log("Test list of tickets fetched")
}

func TestAccess_DeleteTicket(t *testing.T) {
	a := Access{UserName: "username@example.com", Password: "pass", Domain: "acme"}
	err := a.DeleteTicket(1)
	if err.IsError {
		t.Fatalf("error in deleting ticket testing \n%s", err.Message)
	}
	t.Log("Test delete ticket successful")
}
