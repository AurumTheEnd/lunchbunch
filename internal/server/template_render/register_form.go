package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
)

type RegisterFormTemplate struct {
	LayoutTemplate
	TargetUrl    string
	TargetMethod string
	Username     string
	Errors       []string
}

func RenderRegister(w http.ResponseWriter, usernamePrefill string, userData *session.Data, errors []string) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("register_form")
	if parseErr != nil {
		utils.InternalServerError(w, parseErr)
		return
	}

	var data = RegisterFormTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Register",
			UserData:  *userData,
		},
		Username:     usernamePrefill,
		Errors:       errors,
		TargetUrl:    constants.RegisterFormPath,
		TargetMethod: http.MethodPost,
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		utils.InternalServerError(w, renderError)
		return
	}
}
