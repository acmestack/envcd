package util

import (
	"github.com/acmestack/godkits/gox/cryptox/md5x"
)

// EncryptPassword Password generation Policy
//  @param password  string
//  @param salt string
//  @return string
func EncryptPassword(password, salt string) string {
	return md5x.Md5x(password + salt)
}
