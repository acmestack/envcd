package util

import (
	"testing"
)

// EncryptPassword Password generation Policy Test
func TestEncryptPassword(t *testing.T) {
	password := "admin"
	salt := "07929137ab07437c933d6992321ef9fd"

	encryptPassword := EncryptPassword(password, salt)

	// login
	password = "admin"
	encryptPasswordLogin := EncryptPassword(password, salt)
	if encryptPasswordLogin != encryptPassword {
		t.Errorf("TestEncryptPassword() Login failed salt = %v, password %v", salt, password)
	}
}
