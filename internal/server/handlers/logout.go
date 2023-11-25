package handlers

import (
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"net/http"
)

func (app *AppContext) LogoutHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		app.getLogout(w, req)
	default:
		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) getLogout(w http.ResponseWriter, req *http.Request) {
	userData, err := app.LogoutCookie(req, w)
	if err != nil {
		serverError.InternalServerError(w, err)
	}

	template_render.RenderLogoutSuccess(w, userData)
}
