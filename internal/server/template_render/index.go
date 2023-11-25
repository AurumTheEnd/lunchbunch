package template_render

import (
	"lunchbunch/internal/models"
	serverError "lunchbunch/internal/server/error"
	"lunchbunch/web/templates"
	"net/http"
	"time"
)

type IndexTemplate struct {
	LayoutTemplate
	Snapshot models.RestaurantSnapshot
	Today    string
}

func RenderIndex(w http.ResponseWriter, model models.RestaurantSnapshot) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("index")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	var data = IndexTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Menu Voting",
		},
		Snapshot: model,
		Today:    time.Time(model.Date).Format(time.DateOnly),
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
