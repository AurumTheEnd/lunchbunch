package handlers

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetLogout(w http.ResponseWriter, req *http.Request, _ *session.Data) {
	userData, err := app.LogoutCookie(req, w)
	if err != nil {
		utils.InternalServerError(w, err)
	}

	template_render.RenderLogoutSuccess(w, userData)
}
