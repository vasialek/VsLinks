package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetUID returns UID
func GetUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return time.UTC.String()
	}

	return uid.String()
}

// Decode decodes JSON from body of Request
func Decode(rq *http.Request, data interface{}) error {
	ba, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		return err
	}
	defer rq.Body.Close()

	if err = json.Unmarshal(ba, &data); err != nil {
		return err
	}

	return nil
}

// CheckPasswordHash returns true if hashed password is OK
func CheckPasswordHash(passwordHash, passwordClear string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordClear))
	return err == nil
}

// CalculateHash returns hashed string based on password
func CalculateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
