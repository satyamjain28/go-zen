# go-zen

API Wrapper for Zendesk in golang

### Functions implemented

All the function are implemented by the Access struct. 

- GetTicket()
- ListTickets()
- GetUser()


##### Basic Usage
`
go get github.com/satyamjain28/go-zen/pkg
`

```go
package main
import (
    "github.com/satyamjain28/go-zen/pkg"
    "log"
)
func main(){
	a := pkg.Access{UserName: "username@exmaple.com", Password: "password", Domain: "acme"}
	ticket, _ := a.GetTicket(1)
	log.Println(ticket.ID,ticket.Description,ticket.Subject)
}
```
 
