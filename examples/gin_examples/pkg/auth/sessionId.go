package auth

import (
	"log"

	"github.com/gofrs/uuid"
)

func (Authenticator) SessionID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return u2.String()
}
