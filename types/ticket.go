package types

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
