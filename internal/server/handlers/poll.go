package handlers

import (
	"errors"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/database"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/template_render"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"net/http"
	"strconv"
)

func (app *AppContext) GetPoll(w http.ResponseWriter, req *http.Request, userData *session.Data) {
	var id, trimErr = utils.TrimPathToUint(req.URL.Path, constants.PollPathPrefix)
	if trimErr != nil {
		utils.BadRequestError(w, trimErr)
		return
	}

	var result, dbError = database.SelectSnapshotById(app.Db, id)
	if dbError != nil {
		utils.InternalServerError(w, dbError)
		return
	}

	template_render.RenderPoll(w, result, userData)
}

func (app *AppContext) PostPoll(w http.ResponseWriter, req *http.Request, userData *session.Data) {
	if !userData.IsAuthenticated {
		utils.BadRequestError(w, errors.New("cannot cast vote without being logged in"))
		return
	}

	var snapshotId, trimErr = utils.TrimPathToUint(req.URL.Path, constants.PollPathPrefix)
	if trimErr != nil {
		utils.BadRequestError(w, trimErr)
		return
	}

	var parseErr, form = req.ParseForm(), req.Form
	if parseErr != nil {
		utils.BadRequestError(w, parseErr)
		return
	}

	var stringId = form.Get("restaurant_id")
	var shouldCast = form.Has("cast") && !form.Has("uncast")

	var restaurantId, castErr = strconv.ParseUint(stringId, 10, 0)
	if castErr != nil {
		utils.BadRequestError(w, parseErr)
		return
	}

	if dbError := database.UpdateVote(app.Db, uint(restaurantId), userData.UserID, shouldCast); dbError != nil {
		utils.InternalServerError(w, dbError)
	}

	http.Redirect(w, req, constants.PollPath(snapshotId), http.StatusSeeOther)
}
