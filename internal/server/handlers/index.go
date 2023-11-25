package handlers

import (
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/scraping"
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

func (app *AppContext) PostIndex(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	var snapshot, scrapeError = scraping.Scrape(app.C)
	if scrapeError != nil {
		fmt.Println(scrapeError)
		serverError.InternalServerError(w, scrapeError)
		return
	}

	if dbError := database.UpsertScraped(app.Db, snapshot); dbError != nil {
		serverError.InternalServerError(w, dbError)
		return
	}

	template_render.RenderIndex(w, snapshot, userData)
}
