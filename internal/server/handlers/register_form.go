package handlers

import (
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetRegisterForm(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	template_render.RenderRegister(w, "", userData, []string{})
}

func (app *AppContext) PostRegisterForm(w http.ResponseWriter, req *http.Request, userData *session.Data) {
	var registerData = data.RegisterFormData{
		Username:             req.FormValue("username"),
		Password:             req.FormValue("password"),
		PasswordConfirmation: req.FormValue("password_confirmation"),
	}

	if isError, errors := registerData.GatherFormErrors(); isError {
		template_render.RenderRegister(w, registerData.Username, userData, errors)
		return
	}

	var hashedPassword, hashError = utils.HashPassword(registerData.Password)
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
				userData,
				[]string{fmt.Sprintf("Username '%s' is already taken.", registerData.Username)},
			)
		} else {
			serverError.InternalServerError(w, dbError)
		}
		return
	}

	template_render.RenderRegisterSuccess(w, userData)
}
