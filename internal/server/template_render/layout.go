package template_render

import "lunchbunch/internal/session"

type LayoutTemplate struct {
	PageTitle string
	UserData  session.Data
}
