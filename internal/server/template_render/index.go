package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
)

type IndexTemplate struct {
	LayoutTemplate
	Snapshots []models.RestaurantSnapshot
}

func RenderIndex(w http.ResponseWriter, model []models.RestaurantSnapshot, userData *session.Data) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("index")
	if parseErr != nil {
		utils.InternalServerError(w, parseErr)
		return
	}

	var data = IndexTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Menu Voting",
			UserData:  *userData,
		},
		Snapshots: model,
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		utils.InternalServerError(w, renderError)
		return
	}
}
