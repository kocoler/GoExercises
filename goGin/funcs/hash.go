package funcs

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func PasswordHash(password string) (string, string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	salt, err := turnSalt(6)
	if err != nil {
		return "","",err
	}
	h.Write(salt)

	passwordHash := hex.EncodeToString(h.Sum(nil))
	//log.Println(passwordHash)
	return passwordHash,hex.EncodeToString(salt),nil
}

func turnSalt(len int) ([]byte,error) {
	salt := make([]byte, len)
	_,err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return []byte{},err
	}
	return salt,nil
}

func Check(password string, passwordHash string, salt string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false
	}
	h.Write(saltBytes)
	return hex.EncodeToString(h.Sum(nil)) == passwordHash
}