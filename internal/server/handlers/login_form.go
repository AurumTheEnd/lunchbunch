package handlers

import (
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetLoginForm(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	template_render.RenderLoginForm(w, "", userData, []string{})
}

func (app *AppContext) PostLoginForm(w http.ResponseWriter, req *http.Request, userData *session.Data) {
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
			utils.InternalServerError(w, dbError)
		}
		return
	}

	if !utils.IsPasswordHashSame(loginFormData.Password, user.PasswordHash) {
		template_render.RenderLoginForm(w, loginFormData.Username, userData, []string{"Password is invalid."})
		return
	}

	var newUserData, cookieErr = app.LoginCookie(user.Username, user.ID, req, w)
	if cookieErr != nil {
		utils.InternalServerError(w, cookieErr)
		return
	}

	template_render.RenderLoginSuccess(w, newUserData)
}
