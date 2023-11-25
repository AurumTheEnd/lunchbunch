package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	serverError "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/error"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/web/templates"
	"net/http"
)

type LoginFormTemplate struct {
	LayoutTemplate
	TargetUrl    string
	TargetMethod string
	Username     string
	Errors       []string
}

func RenderLoginForm(w http.ResponseWriter, username string, userData *session.Data, errors []string) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("login_form")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	var data = LoginFormTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Login",
			UserData:  *userData,
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
