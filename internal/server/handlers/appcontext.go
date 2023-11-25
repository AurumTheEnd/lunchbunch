package handlers

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/gorilla/sessions"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gorm.io/gorm"
	"net/http"
)

type AppContext struct {
	Db          *gorm.DB
	C           *colly.Collector
	CookieStore *sessions.CookieStore
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
type AppContextHandler = func(w http.ResponseWriter, req *http.Request, userData *session.Data)

func (app *AppContext) ReusableHandler(getHandler AppContextHandler, postHandler AppContextHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		userData, sessionErr := app.UserData(req)
		if sessionErr != nil {
			serverError.InternalServerError(w, sessionErr)
		}

		switch req.Method {
		case http.MethodGet:
			if getHandler == nil {
				break
			}
			getHandler(w, req, userData)
		case http.MethodPost:
			if postHandler == nil {
				break
			}
			postHandler(w, req, userData)
		}

		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) DisallowSubtreeWrapper(permittedPath string) Middleware {
	return func(nextMiddleWare http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path != permittedPath {
				http.NotFound(w, req)
				return
			}

			nextMiddleWare(w, req)
		}
	}
}

func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
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

func (app *AppContext) UserData(req *http.Request) (*session.Data, error) {
	var s, getError = app.CookieStore.Get(req, session.CookieName)
	if getError != nil {
		return nil, getError
	}

	var userData, ok = s.Values[session.AuthenticationStoreKey].(session.Data)
	if !ok {
		return nil, errors.New("session value has invalid type")
	}

	return &userData, nil
}
