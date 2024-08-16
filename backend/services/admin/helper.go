package admin

import "golang.org/x/crypto/bcrypt"

func hash(v string) (string, error) {
	password := []byte(v)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
