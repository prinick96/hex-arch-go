package helpers_test

import (
	"hex-arch-go/internal/helpers"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	testCases := []struct {
		Name string
		UUID string
	}{
		{
			Name: "It should be a valid UUID",
			UUID: helpers.UUID(),
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			uui, err := uuid.Parse(tc.UUID)

			if err != nil {
				t.Errorf("The UUID `%s` is invalid", uui.String())
			}
		})
	}

}
