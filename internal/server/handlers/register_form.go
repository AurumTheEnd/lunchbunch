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

func (app *AppContext) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		app.getRegisterForm(w, req)
	case http.MethodPost:
		app.postRegisterForm(w, req)
	default:
		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) getRegisterForm(w http.ResponseWriter, _ *http.Request) {
	template_render.RenderRegister(w, "", []string{})
}

func (app *AppContext) postRegisterForm(w http.ResponseWriter, req *http.Request) {
	var registerData = data.RegisterFormData{
		Username:             req.FormValue("username"),
		Password:             req.FormValue("password"),
		PasswordConfirmation: req.FormValue("password_confirmation"),
	}

	if isError, errors := registerData.GatherFormErrors(); isError {
		template_render.RenderRegister(w, registerData.Username, errors)
		return
	}

	var hashedPassword, hashError = auth.HashPassword(registerData.Password)
	if hashError != nil {
		serverError.InternalServerError(w, hashError)
		return
	}

	var dbError = database.CreateUser(app.Db, registerData.Username, hashedPassword)
	if dbError != nil {
		if database.IsUniqueViolation(dbError) {
			template_render.RenderRegister(
				w,
				registerData.Username,
				[]string{fmt.Sprintf("Username '%s' is already taken.", registerData.Username)},
			)
		} else {
			serverError.InternalServerError(w, dbError)
		}
		return
	}

	data, sessionErr := app.LogoutCookie(req, w)
	if sessionErr != nil {
		serverError.InternalServerError(w, sessionErr)
	}

	template_render.RenderRegisterSuccess(w, data)
}
