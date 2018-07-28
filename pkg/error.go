package pkg

// ErrorResponse standard wrapper for error in go-zen
type ErrorResponse struct {
	Message string
	Code    int
	IsError bool
}
