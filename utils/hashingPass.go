package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pass string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)

}
