package template_render

import (
	"lunchbunch/internal/server/constants"
	serverError "lunchbunch/internal/server/error"
	"lunchbunch/web/templates"
	"net/http"
)

type LoginFormTemplate struct {
	LayoutTemplate
	TargetUrl    string
	TargetMethod string
	Username     string
	Errors       []string
}

func RenderLoginForm(w http.ResponseWriter, username string, errors []string) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("login_form")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	var data = LoginFormTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Login",
		},
		Errors:       errors,
		Username:     username,
		TargetUrl:    constants.LoginFormPath,
		TargetMethod: "POST",
	}

	if errors != nil && len(errors) != 0 {
		w.WriteHeader(http.StatusForbidden)
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
