package template

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	templates = make(map[string]*template.Template)
)

// InitTemplate .
func InitTemplate(templatePath string) error {
	layouts, err := filepath.Glob(templatePath + "layout/*.html")
	if err != nil {
		// log.Error.Println("initialize template", err)
		return err
	}

	for _, layout := range layouts {
		t := template.Must(template.ParseFiles(layout))
		if err != nil {
			// log.Error.Println("parse template", err)
			return err
		}
		t.Delims("[[", "]]")
		templates[filepath.Base(layout)] = t
	}
	// log.Info.Println("initialize template success")
	return nil
}

// RenderTemplate .
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, correct := templates[name]
	if !correct {
		// log.Error.Println("can not find template file: " + name, err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}

// RenderNotFound .
func RenderNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
