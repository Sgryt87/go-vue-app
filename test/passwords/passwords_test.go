package passwords

import (
	"CODE/goVue/src/system/passwords"
	"log"
	"testing"
)

func init() {
	log.Println("Testing password")
}

func TestBasicLog(t *testing.T) {
	pass := "TEST"
	hash, err := passwords.Encrypt(pass)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(hash)
	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Passwords not matching")
	}
	log.Println("Successfully tested hashing")
}
