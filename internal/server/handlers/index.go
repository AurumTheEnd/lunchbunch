package handlers

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetIndex(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	var result, dbError = database.SelectTodaysSnapshot(app.Db)
	if dbError != nil {
		serverError.InternalServerError(w, dbError)
	}

	template_render.RenderIndex(w, result, userData)
}
