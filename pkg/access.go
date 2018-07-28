package pkg

// Access struct defines the credentials for the user
type Access struct {
	UserName string
	Password string
	Domain   string // Initialize domain as example if your zendesk endpoint is example.zendesk.com
}
