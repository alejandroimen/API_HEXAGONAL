package core

import "golang.org/x/crypto/bcrypt"

func HashAttribute(attribute string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(attribute), bcrypt.DefaultCost)
	if err != nil {
		return "error"
	}
	return string(hashed)
}
