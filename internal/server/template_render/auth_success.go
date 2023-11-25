package template_render

import (
	serverError "lunchbunch/internal/server/error"
	"lunchbunch/internal/session"
	"lunchbunch/web/templates"
	"net/http"
)

type AuthSuccessTemplate struct {
	LayoutTemplate
	ButtonUrl  string
	ButtonText string
	Message    string
}

func RenderLogoutSuccess(w http.ResponseWriter, userData *session.Data) {
	var data = AuthSuccessTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Logged out",
			UserData:  *userData,
		},
		ButtonUrl:  "/login",
		ButtonText: "Log in again",
		Message:    "Logged out successfully!",
	}

	renderAuthSuccess(w, data)
}

func RenderRegisterSuccess(w http.ResponseWriter, userData *session.Data) {
	var data = AuthSuccessTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Registered",
			UserData:  *userData,
		},
		ButtonUrl:  "/login",
		ButtonText: "Log in now!",
		Message:    "Registered successfully!",
	}

	renderAuthSuccess(w, data)
}

func RenderLoginSuccess(w http.ResponseWriter, userData *session.Data) {
	var data = AuthSuccessTemplate{
		LayoutTemplate: LayoutTemplate{
			PageTitle: "Logged in",
			UserData:  *userData,
		},
		ButtonUrl:  "",
		ButtonText: "",
		Message:    "Logged in successfully!",
	}

	renderAuthSuccess(w, data)
}

func renderAuthSuccess(w http.ResponseWriter, data AuthSuccessTemplate) {
	var parsedTemplate, parseErr = templates.ParseTemplateWithLayout("auth_success")
	if parseErr != nil {
		serverError.InternalServerError(w, parseErr)
		return
	}

	if renderError := parsedTemplate.Execute(w, data); renderError != nil {
		serverError.InternalServerError(w, renderError)
		return
	}
}
