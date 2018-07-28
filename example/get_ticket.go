package example

import (
	"log"

	"github.com/satyamjain28/go-zen/pkg"
)

func GetTicket() {
	a := pkg.Access{UserName: "username@exmaple.com", Password: "password", Domain: "example"}
	ticket, err := a.GetTicket(11)
	if !err.IsError {
		log.Println(ticket.ID, ticket.Description, ticket.Subject)
	}
}
