package validators

import (
	"regexp"

	"github.com/google/uuid"
)


func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func IsValidName(name string) bool {
	length := len(name)
	return 3 <= length && length <= 200
}

func IsValidUuid(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}