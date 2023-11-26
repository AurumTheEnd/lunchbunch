package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	formData "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
)

type NewPollFormTemplate struct {
	LayoutTemplate
	formData.NewPollFormDataToClient
	TargetUrl    string
	TargetMethod string
	Errors       []string
}

func RenderNewPollForm(w http.ResponseWriter, formData formData.NewPollFormDataToClient, userData *session.Data, errors []string) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("new_poll_form")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	// display checkboxes
	// Low priority: get from Session that never expires: list of ids to prefill

	var data = NewPollFormTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "New poll",
			UserData:  *userData,
		},
		NewPollFormDataToClient: formData,
		TargetUrl:               constants.NewPollFormPath,
		TargetMethod:            http.MethodPost,
		Errors:                  errors,
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
