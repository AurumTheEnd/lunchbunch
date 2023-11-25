package handlers

import (
	"fmt"
	"lunchbunch/internal/database"
	"lunchbunch/internal/server/auth"
	"lunchbunch/internal/server/data"
	serverError "lunchbunch/internal/server/error"
	"lunchbunch/internal/server/template_render"
	"net/http"
)

func (app *AppContext) LoginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		app.getLoginForm(w, req)
	case http.MethodPost:
		app.postLoginForm(w, req)
	default:
		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) getLoginForm(w http.ResponseWriter, _ *http.Request) {
	template_render.RenderLoginForm(w, "", []string{})
}

func (app *AppContext) postLoginForm(w http.ResponseWriter, req *http.Request) {
	var loginFormData = data.LoginFormData{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if isError, errors := loginFormData.GatherFormErrors(); isError {
		template_render.RenderLoginForm(w, loginFormData.Username, errors)
		return
	}

	var user, dbError = database.GetUser(app.Db, loginFormData.Username)
	if dbError != nil {
		if database.IsRecordNotFound(dbError) {
			template_render.RenderLoginForm(
				w,
				loginFormData.Username,
				[]string{fmt.Sprintf("User '%s' doesn't exist.", loginFormData.Username)},
			)
		} else {
			serverError.InternalServerError(w, dbError)
		}
		return
	}

	if !auth.IsPasswordHashSame(loginFormData.Password, user.PasswordHash) {
		template_render.RenderLoginForm(w, loginFormData.Username, []string{"Password is invalid."})
		return
	}

	var userData, cookieErr = app.LoginCookie(user.Username, req, w)
	if cookieErr != nil {
		serverError.InternalServerError(w, cookieErr)
		return
	}

	template_render.RenderLoginSuccess(w, userData)
}
