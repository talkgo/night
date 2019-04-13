package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	var envs = []struct {
		key   string
		value string
	}{
		{"PORT", "testport"},
		{"ENV", "testenv"},
		{"PG_HOST", "testhost"},
		{"PG_PORT", "testport"},
		{"PG_USER", "testuser"},
		{"PG_PASSWORD", "testpass"},
		{"PG_DB_NAME", "testname"},
		{"LOGFILE", "backend.log"},
	}

	var testCases = []struct {
		name     string
		unset    bool
		expected Config
	}{
		{"set", false,
			Config{
				"testport",
				"testenv",
				"testhost",
				"testport",
				"testuser",
				"testpass",
				"testname",
				"backend.log",
			}},
		{"unset", true,
			Config{
				"8080",
				"development",
				"localhost",
				"5432",
				"postgres",
				"",
				"ginexamples",
				"",
			}},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			for _, e := range envs {
				os.Unsetenv(e.key)
				if !v.unset {
					os.Setenv(e.key, e.value)
				}
			}

			c := GetConfig()
			cExpected := v.expected
			assert.Equal(t, cExpected, c, "config does not match expected one")
		})
	}
}
