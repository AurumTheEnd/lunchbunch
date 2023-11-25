package handlers

import (
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/auth"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) LoginHandler(w http.ResponseWriter, req *http.Request) {
	userData, sessionErr := app.UserData(req)
	if sessionErr != nil {
		serverError.InternalServerError(w, sessionErr)
	}

	switch req.Method {
	case http.MethodGet:
		app.getLoginForm(w, req, userData)
	case http.MethodPost:
		app.postLoginForm(w, req, userData)
	default:
		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) getLoginForm(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	template_render.RenderLoginForm(w, "", userData, []string{})
}

func (app *AppContext) postLoginForm(w http.ResponseWriter, req *http.Request, userData *session.Data) {
	var loginFormData = data.LoginFormData{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if isError, errors := loginFormData.GatherFormErrors(); isError {
		template_render.RenderLoginForm(w, loginFormData.Username, userData, errors)
		return
	}

	var user, dbError = database.GetUser(app.Db, loginFormData.Username)
	if dbError != nil {
		if database.IsRecordNotFound(dbError) {
			template_render.RenderLoginForm(
				w,
				loginFormData.Username,
				userData,
				[]string{fmt.Sprintf("User '%s' doesn't exist.", loginFormData.Username)},
			)
		} else {
			serverError.InternalServerError(w, dbError)
		}
		return
	}

	if !auth.IsPasswordHashSame(loginFormData.Password, user.PasswordHash) {
		template_render.RenderLoginForm(w, loginFormData.Username, userData, []string{"Password is invalid."})
		return
	}

	var newUserData, cookieErr = app.LoginCookie(user.Username, req, w)
	if cookieErr != nil {
		serverError.InternalServerError(w, cookieErr)
		return
	}

	template_render.RenderLoginSuccess(w, newUserData)
}
