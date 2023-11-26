package template_render

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/constants"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/utils"
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
		utils.InternalServerError(w, parseErr)
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
		TargetMethod: http.MethodPost,
	}

	if errors != nil && len(errors) != 0 {
		w.WriteHeader(http.StatusForbidden)
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		utils.InternalServerError(w, renderError)
		return
	}
}
