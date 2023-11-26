package server

import (
	"errors"
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/handlers"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type AppContextHandler = func(w http.ResponseWriter, req *http.Request, userData *session.Data)

func ReusableHandler(appContext *handlers.AppContext, allowUnauthorized bool, getHandler AppContextHandler, postHandler AppContextHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userData = appContext.UserData(req)

		if !allowUnauthorized && !userData.IsAuthenticated {
			utils.UnauthorizedError(w, errors.New(fmt.Sprintf("Unauthorized acces to %s", req.URL)))
			return
		}

		switch req.Method {
		case http.MethodGet:
			if getHandler == nil {
				break
			}
			getHandler(w, req, userData)
			return
		case http.MethodPost:
			if postHandler == nil {
				break
			}
			postHandler(w, req, userData)
			return
		}

		utils.MethodNotAllowed(w, req.Method)
	}
}

func DisallowSubtreeWrapper(permittedPath string) Middleware {
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

func AllowIdSuffixOnly(prefix string) Middleware {
	return func(nextMiddleWare http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			var path = req.URL.Path
			if !strings.HasPrefix(path, prefix) {
				utils.InternalServerError(w, errors.New(fmt.Sprintf("request path '%s' doesn't start with prefix '%s' to strip", path, prefix)))
				return
			}

			var _, convertErr = utils.TrimPathToUint(path, prefix)
			if convertErr != nil {
				utils.BadRequestError(w, convertErr)
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
