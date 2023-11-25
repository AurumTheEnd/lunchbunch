package handlers

import (
	"github.com/gocolly/colly"
	"github.com/gorilla/sessions"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gorm.io/gorm"
	"net/http"
)

type AppContext struct {
	Db          *gorm.DB
	C           *colly.Collector
	CookieStore *sessions.CookieStore
}

func (app *AppContext) LoginCookie(username string, req *http.Request, w http.ResponseWriter) (*session.Data, error) {
	var s, getError = app.CookieStore.Get(req, session.CookieName)
	if getError != nil {
		return nil, getError
	}

	var data = session.Data{
		Authenticated: true,
		Username:      username,
	}

	s.Values[session.AuthenticationStoreKey] = data

	if saveErr := s.Save(req, w); saveErr != nil {
		return nil, saveErr
	}

	return &data, nil
}

func (app *AppContext) LogoutCookie(req *http.Request, w http.ResponseWriter) (*session.Data, error) {
	var s, getError = app.CookieStore.Get(req, session.CookieName)
	if getError != nil {
		return nil, getError
	}

	var data = session.Data{
		Authenticated: false,
	}

	s.Values[session.AuthenticationStoreKey] = data

	if saveErr := s.Save(req, w); saveErr != nil {
		return nil, saveErr
	}

	return &data, nil
}

func (app *AppContext) IsAuthenticated(req *http.Request) (bool, error) {
	var s, getError = app.CookieStore.Get(req, session.CookieName)
	if getError != nil {
		return false, getError
	}

	var userData, ok = s.Values[session.AuthenticationStoreKey].(session.Data)

	return !ok || !userData.Authenticated, nil
}

func (app *AppContext) UserData(req *http.Request) *session.Data {
	var s, _ = app.CookieStore.Get(req, session.CookieName)
	// ignoring error because Get generates a new session
	// if getError != nil {
	// 	 return nil, getError
	// }

	var userData, ok = s.Values[session.AuthenticationStoreKey].(session.Data)
	if !ok {
		return &session.Data{}
	}

	return &userData
}
