package helpers

import (
	"github.com/google/uuid"
)

// Create a UUID as string
func UUID() string {
	return uuid.NewString()
}
