package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptPassword Password generation Policy
//  @param password  string
//  @param salt string
// return string
func EncryptPassword(password, salt string) string {
	h := md5.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}
