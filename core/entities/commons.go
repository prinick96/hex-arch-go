package entities

// Commons types
type Response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}
