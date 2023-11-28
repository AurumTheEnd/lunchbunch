package handlers

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetIndex(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	var result, dbError = database.SelectTodaysSnapshots(app.Db)
	if dbError != nil {
		utils.InternalServerError(w, dbError)
	}

	template_render.RenderIndex(w, result, userData)
}
