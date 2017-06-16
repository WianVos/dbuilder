package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Template struct {
	RelativeName string
	Tmpl         *template.Template
}

type Templates []Template

func NewTemplate(r string, t *template.Template) Template {
	return Template{
		RelativeName: r,
		Tmpl:         t,
	}
}

func (t *Template) Write(d string, data interface{}) error {
	f := d + t.RelativeName
	fmt.Println(f)
	if _, err := os.Stat(filepath.Dir(f)); err != nil {
		err := os.MkdirAll(filepath.Dir(f), os.FileMode(fileMode))
		if err != nil {
			return err
		}
	}
	iow, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0770)
	if err != nil {
		fmt.Println(err)
		return err
	}

	t.Tmpl.Execute(iow, data)

	return nil
}
