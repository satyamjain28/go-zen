package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

// TicketRes response structure for ticket
type TicketRes struct {
	Ticket Ticket `json:"ticket"`
}

// TicketListRes response structure for ticket list
type TicketListRes struct {
	Tickets []Ticket `json:"tickets"`
	Count   int      `json:"count"`
}

// Ticket structure for zendesk ticket
type Ticket struct {
	ID             int    `json:"id,omitempty"`
	URL            string `json:"url,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	Type           string `json:"type,omitempty"`
	Subject        string `json:"subject,omitempty"`
	RawSubject     string `json:"raw_subject,omitempty"`
	Description    string `json:"description,omitempty"`
	Priority       string `json:"priority,omitempty"`
	Status         string `json:"status,omitempty"`
	Recipient      string `json:"recipient,omitempty"`
	RequesterID    int    `json:"requester_id,omitempty"`
	SubmitterID    int    `json:"submitter_id,omitempty"`
	AssigneeID     int    `json:"assignee_id,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
	GroupID        int    `json:"group_id,omitempty"`
	CustomFields   []struct {
		ID    int         `json:"id,omitempty"`
		Value interface{} `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
	TicketFormID int    `json:"ticket_form_id,omitempty"`
	BrandID      int    `json:"brand_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// GetTicket gets the ticket for the given ticket id
func (a Access) GetTicket(ticketID int) (*Ticket, ErrorResponse) {
	var ticketRes TicketRes
	resp, errRes := NewRequest(a, fmt.Sprintf("tickets/%d.json", ticketID), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("Error in fetching ticket with id %d", ticketID)
		return &ticketRes.Ticket, errRes
	}
	err := json.Unmarshal(resp, &ticketRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response to the ticketRes body \n %v ", err)
		return &ticketRes.Ticket, ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}
	return &ticketRes.Ticket, ErrorResponse{IsError: false}
}

// ListTickets lists the tickets on the given page id
func (a Access) ListTickets(page int) (*[]Ticket, ErrorResponse) {
	var ticketRes TicketListRes
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
		return &ticketRes.Tickets, ErrorResponse{Message: err.Error(), Code: 400, IsError: true}
	}
	return &ticketRes.Tickets, ErrorResponse{IsError: false}
}

// ListTicketsByExternalID lists the tickets sorted by the given external id
func (a Access) ListTicketsByExternalID(externalID string) (*[]Ticket, ErrorResponse) {
	var ticketListRes TicketListRes
	resp, errRes := NewRequest(a, fmt.Sprintf("tickets.json?external_id=%s", externalID), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("error in fetching the list of tickets")
		return &ticketListRes.Tickets, errRes
	}
	err := json.Unmarshal(resp, &ticketListRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response body into ticketListRes struct")
		return &ticketListRes.Tickets, ErrorResponse{IsError: true, Code: 400, Message: err.Error()}
	}
	return &ticketListRes.Tickets, ErrorResponse{IsError: false}
}

// DeleteTicket deletes the ticket with given Id
func (a Access) DeleteTicket(ticketID int) ErrorResponse {
	_, errRes := NewRequest(a, fmt.Sprintf("tickets/%d.json", ticketID), "DELETE", nil, 204)
	if errRes.IsError {
		log.Printf("Error in deleting the ticket")
		return errRes
	}
	return ErrorResponse{IsError: false}
}

// ListTicketsByID provides details for the given ticket IDs
func (a Access) ListTicketsByID(ticketIDs []int) (*[]Ticket, ErrorResponse) {
	var ticketListRes TicketListRes
	ticketIDString := ""
	for _, v := range ticketIDs {
		if ticketIDString == "" {
			ticketIDString = fmt.Sprintf("%d", v)
		} else {
			ticketIDString += fmt.Sprintf(",%d", v)
		}
	}
	resp, errRes := NewRequest(a, fmt.Sprintf("tickets/show_many.json?ids=%s", ticketIDString), "GET", nil, 200)
	if errRes.IsError {
		log.Printf("error in fetching the list of given tickets")
		return &ticketListRes.Tickets, errRes
	}
	err := json.Unmarshal(resp, &ticketListRes)
	if err != nil {
		log.Printf("Error in unmarshalling the response body into ticketListRes struct")
		return &ticketListRes.Tickets, ErrorResponse{IsError: true, Code: 400, Message: err.Error()}
	}
	return &ticketListRes.Tickets, ErrorResponse{IsError: false}
}
