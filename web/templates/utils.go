package templates

import (
	"html/template"
	"os"
	"strings"
)

const pathToTemplatesFromRoot = "web" + string(os.PathSeparator) + "templates"
const layoutTemplateName = "base"

func maybeAppendGohtml(s string) string {
	if strings.HasSuffix(s, ".gohtml") {
		return s
	}

	return s + ".gohtml"
}

func ParseTemplateWithLayout(templateName string) (*template.Template, error) {
	var templatePath = pathToTemplatesFromRoot + string(os.PathSeparator) + maybeAppendGohtml(templateName)
	var layoutPath = pathToTemplatesFromRoot + string(os.PathSeparator) + maybeAppendGohtml(layoutTemplateName)

	// layout must come last
	return template.ParseFiles(templatePath, layoutPath)
}
