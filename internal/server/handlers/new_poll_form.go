package handlers

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/scraping"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
)

func (app *AppContext) GetNewPollForm(w http.ResponseWriter, _ *http.Request, userData *session.Data) {
	// scrape new data!
	var snapshot, scrapeError = scraping.Scrape(app.C)
	if scrapeError != nil {
		utils.InternalServerError(w, scrapeError)
		return
	}

	snapshot.CreatorID = userData.UserID
	snapshot.HasPollStarted = false

	// save scraped new data to db
	if dbError := database.CreateScraped(app.Db, &snapshot); dbError != nil {
		utils.InternalServerError(w, dbError)
		return
	}

	var formData = data.FromSnapshot(&snapshot)

	template_render.RenderNewPollForm(w, formData, userData, []string{})
}

func (app *AppContext) PostNewPollForm(w http.ResponseWriter, req *http.Request, userData *session.Data) {
	// Low priority: update Session that never expires: list of ids to prefill

	// parse form with checkboxes
	var formResult, err = data.FromRequest(req)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// validate form data: disallow 0 checkboxes
	if len(formResult.Checked) == 0 {
		var snapshotAgain, dbError = database.SelectSnapshotWith(app.Db, formResult.SnapshotID)
		if dbError != nil {
			utils.InternalServerError(w, dbError)
			return
		}

		var formData = data.FromSnapshot(&snapshotAgain)
		template_render.RenderNewPollForm(w, formData, userData, []string{"You must select at least one restaurant."})
		return
	}

	// update DB: restaurant.voted_on (for later filtering to display in poll)
	var dbError = database.UpdateVotedOn(app.Db, formResult)
	if dbError != nil {
		utils.InternalServerError(w, dbError)
		return
	}

	// redirect to /poll/{id} of restaurant_snapshot.id
	http.Redirect(w, req, constants.PollPath(formResult.SnapshotID), http.StatusFound)
}
