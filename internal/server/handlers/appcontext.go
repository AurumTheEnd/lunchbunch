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
	var s, _ = app.CookieStore.Get(req, session.AuthCookieName)
	// ignoring error because Get generates a new session
	// if getError != nil {
	// 	 return nil, getError
	// }

	var data = &session.Data{
		IsAuthenticated: true,
		Username:        username,
	}

	s.Values[session.AuthenticationStoreKey] = data

	if saveErr := s.Save(req, w); saveErr != nil {
		return &session.Data{IsAuthenticated: false}, saveErr
	}

	return data, nil
}

func (app *AppContext) LogoutCookie(req *http.Request, w http.ResponseWriter) (*session.Data, error) {
	var s, _ = app.CookieStore.Get(req, session.AuthCookieName)
	// ignoring error because Get generates a new session
	// if getError != nil {
	// 	 return nil, getError
	// }

	var data = &session.Data{IsAuthenticated: false}

	s.Values[session.AuthenticationStoreKey] = data
	s.Options.MaxAge = -1

	if saveErr := s.Save(req, w); saveErr != nil {
		return data, saveErr
	}

	return data, nil
}

func (app *AppContext) IsAuthenticated(req *http.Request) (bool, error) {
	var s, _ = app.CookieStore.Get(req, session.AuthCookieName)
	// ignoring error because Get generates a new session
	// if getError != nil {
	// 	 return nil, getError
	// }

	var userData, ok = s.Values[session.AuthenticationStoreKey].(session.Data)

	return !ok || !userData.IsAuthenticated, nil
}

func (app *AppContext) UserData(req *http.Request) *session.Data {
	var s, _ = app.CookieStore.Get(req, session.AuthCookieName)
	// ignoring error because Get generates a new session
	// if getError != nil {
	// 	 return nil, getError
	// }

	var raw = s.Values[session.AuthenticationStoreKey]

	var userData, ok = raw.(*session.Data)
	if !ok {
		// data is mangled, maybe using old cookie format, lets revoke login
		return &session.Data{IsAuthenticated: false}
	}

	return userData
}
