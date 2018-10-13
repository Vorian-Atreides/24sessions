package backend

// Error represent an API error for the end user
type Error struct {
	// Reason or source of the error
	Reason string `json:"reason"`
}

// NewError instantiate a new Error
func NewError(reason string) *Error {
	return &Error{Reason: reason}
}
