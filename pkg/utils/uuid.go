package utils

import (
	"github.com/oklog/ulid/v2"
)

// GenerateULID generates a new ULID
func GenerateULID() string {
	entropy := ulid.Monotonic(nil, 0)
	return ulid.MustNew(ulid.Now(), entropy).String()
}
