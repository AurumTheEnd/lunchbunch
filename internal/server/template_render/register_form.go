package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
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

func RenderRegister(w http.ResponseWriter, usernamePrefill string, errors []string) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("register_form")
	if parseErr != nil {
		error.InternalServerError(w, parseErr)
		return
	}

	var data = RegisterFormTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Register",
		},
		Username:     usernamePrefill,
		Errors:       errors,
		TargetUrl:    constants.RegisterFormPath,
		TargetMethod: "POST",
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		error.InternalServerError(w, renderError)
		return
	}
}
