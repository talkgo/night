package auth

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_SessionID(t *testing.T) {
	t.Parallel()
	authenticator := Authenticator{}
	sessionID := authenticator.SessionID()

	id, err := uuid.FromString(sessionID)
	assert.Nil(t, err, "should not cause error")
	assert.Equal(t, byte(4), id.Version(), "should create UUID v4")
}
