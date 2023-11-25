package session

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func generateKey(size uint) (key []byte, err error) {
	key = make([]byte, size)

	_, err = rand.Read(key)

	return
}

func CreateSessionStore() (store *sessions.CookieStore, err error) {
	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	var auth64 = myEnv["SESSION_AUTH_KEY"]
	var encrypt64 = myEnv["SESSION_ENCRYPT_KEY"]

	var auth, decodeAuthErr = base64.StdEncoding.DecodeString(auth64)
	if decodeAuthErr != nil {
		return nil, decodeAuthErr
	}
	var encrypt, decodeEncryptErr = base64.StdEncoding.DecodeString(encrypt64)
	if decodeEncryptErr != nil {
		return nil, decodeEncryptErr
	}

	store = sessions.NewCookieStore(
		auth,
		encrypt,
	)

	gob.Register(Data{})

	return
}
