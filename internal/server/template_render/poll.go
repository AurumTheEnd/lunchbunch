package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
)

type PollTemplate struct {
	LayoutTemplate
	Snapshot models.RestaurantSnapshot
}

func RenderPoll(w http.ResponseWriter, model models.RestaurantSnapshot, userData *session.Data) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("poll")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	var data = PollTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Poll",
			UserData:  *userData,
		},
		Snapshot: model,
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
