package auth

import (
	"perfume/packages/exception"

	"golang.org/x/crypto/bcrypt"
)

// PasswordEncode is the method to encode password
func PasswordEncode(password string) string {
	encodePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	recoder.Write(err)
	return string(encodePassword)
}

// PasswordCheck id the method to check password match
func PasswordCheck(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return recoder.Write(err)
}
