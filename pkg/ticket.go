package pkg

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/satyamjain28/go-zen/types"
)

// GetTicket gets the ticket for the given ticket id
func (a Access) GetTicket(ticketID int) (*types.Ticket, types.ErrorResponse) {
	var ticketRes types.TicketRes
	resp, errRes := NewRequest(a, fmt.Sprintf("tickets/%d.json", ticketID), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("Error in fetching ticket with id %d", ticketID)
		return &ticketRes.Ticket, errRes
	}
	err := json.Unmarshal(resp, &ticketRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response to the ticketRes body \n %v ", err)
		return &ticketRes.Ticket, types.ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}
	return &ticketRes.Ticket, types.ErrorResponse{IsError: false}
}

// ListTickets lists the tickets on the given page id
func (a Access) ListTickets(page int) (*[]types.Ticket, types.ErrorResponse) {
	var ticketRes types.TicketListRes
	if page == 0 {
		page = 1
	}
	resp, errRes := NewRequest(a, fmt.Sprintf("tickets.json?page=%d", page), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("Error in fetching list of tickets ")
		return &ticketRes.Tickets, errRes
	}
	err := json.Unmarshal(resp, &ticketRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response to the tickets res body")
		return &ticketRes.Tickets, types.ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}
	return &ticketRes.Tickets, types.ErrorResponse{IsError: false}
}
