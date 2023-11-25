package template_render

import "gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/session"

type LayoutTemplate struct {
	PageTitle string
	UserData  session.Data
}
