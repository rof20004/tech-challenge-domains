package utils

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsEmailValido(email string) bool {
	if ok := emailRegex.MatchString(email); !ok {
		return false
	}

	return true
}
