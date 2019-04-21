package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticationProvider_HashAndSalt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name  string
		input string
	}{
		{"expected", "password"},
		{"empty", ""},
	}

	authenticator := Authenticator{}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			hash, err := authenticator.Hash(v.input)
			assert.Nil(t, err, "should not cause error")
			assert.NotEmptyf(t, hash, "hashes do not match")

			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(v.input))
			assert.Nil(t, err, "passwords do not match")
		})
	}
}

func TestAuthenticator_CompareWithHash(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name      string
		input     string
		reference string
	}{
		{"expected", "password", "password"},
		{"empty", "", ""},
		{"error", "password", "drowssap"},
	}

	authenticator := Authenticator{}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			hash, err := authenticator.Hash(v.reference)
			assert.Nil(t, err, "should not cause error")
			assert.NotEmptyf(t, hash, "hashes do not match")

			err = authenticator.CompareHash(hash, v.input)

			if v.input != v.reference {
				assert.Error(t, err, "should cause an error")
				return
			}
			assert.Nil(t, err, "should not cause error")
		})
	}
}
