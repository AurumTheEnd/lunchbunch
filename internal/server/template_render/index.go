package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
	"time"
)

type IndexTemplate struct {
	LayoutTemplate
	Snapshot models.RestaurantSnapshot
	Today    string
}

func RenderIndex(w http.ResponseWriter, model models.RestaurantSnapshot, userData *session.Data) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("index")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	var data = IndexTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Menu Voting",
			UserData:  *userData,
		},
		Snapshot: model,
		Today:    time.Time(model.Date).Format(time.DateOnly),
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
