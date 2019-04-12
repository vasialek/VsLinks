package helpers

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// GetUID returns UID
func GetUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return time.UTC.String()
	}

	return uid.String()
}
