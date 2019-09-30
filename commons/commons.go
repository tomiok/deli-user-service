package commons

import "github.com/google/uuid"

func StringUUID() string {
	return uuid.New().String()
}
