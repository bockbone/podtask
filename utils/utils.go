package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	password := []byte(p)
	hpassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hpassword)
}
