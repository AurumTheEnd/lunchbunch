package session

import (
	"net/http"
)

type CookieStoreOperation interface {
	LoginCookie(username string, req *http.Request, w http.ResponseWriter) (*Data, error)
	LogoutCookie(req *http.Request, w http.ResponseWriter) (*Data, error)
	IsAuthenticated(req *http.Request) (bool, error)
	UserData(req *http.Request) (*Data, error)
}
