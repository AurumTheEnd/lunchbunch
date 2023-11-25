package handlers

import (
	"fmt"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/scraping"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) IndexHandler() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != constants.IndexPath { // Check path here
			http.NotFound(w, req)
			return
		}

		app.indexHandler(w, req)
	}
}

func (app *AppContext) indexHandler(w http.ResponseWriter, req *http.Request) {
	userData, sessionErr := app.UserData(req)
	if sessionErr != nil {
		serverError.InternalServerError(w, sessionErr)
	}

	switch req.Method {
	case http.MethodGet:
		app.getIndex(w, req, userData)
	case http.MethodPost:
		app.postIndex(w, req, userData)
	default:
		serverError.MethodNotAllowed(w, req.Method)
	}
}

func (app *AppContext) getIndex(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	var result, dbError = database.SelectTodaysSnapshot(app.Db)
	if dbError != nil {
		serverError.InternalServerError(w, dbError)
	}

	template_render.RenderIndex(w, result, userData)
}

func (app *AppContext) postIndex(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
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
