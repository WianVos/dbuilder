package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/satori/go.uuid"
	"github.com/wianvos/dbuilder/pkg/config"
)

const (
	requirementsFile = "dbuilder.req"
	templateDir      = "/templates"
	templateDefaults = templateDir + "/default"
	buildDir         = "/.build"
	fileMode         = 0770
)

type Repo struct {
	Path        string
	BuildDir    string
	Environment string
	Templates   Templates
	Config      config.Config
}

// returns a new Repo object
func NewRepo(p string) (Repo, error) {

	r := Repo{
		Path: p,
	}

	r.FindTemplates()

	bdir, err := getBuildDir(p)
	if err != nil {
		return Repo{}, err
	}

	r.BuildDir = bdir
	return r, nil

}

//templateSearchPath returns the constructed search path
func (r Repo) templateSearchPath() []string {

	s := []string{}

	if r.Environment != "" {
		envString := r.Path + templateDir + " / " + r.Environment
		s = append(s, envString)
	}

	s = append(s, r.Path+templateDefaults)

	return s
}

func (r *Repo) FindTemplates() {
	tfs := Templates{}

	for _, sp := range r.templateSearchPath() {
		list := []string{}
		err := filepath.Walk(sp, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			list = append(list, filepath.Clean(path))

			return nil
		})
		if err != nil {
			fmt.Printf("walk error [%v]\n", err)
		}

		for _, f := range list {
			relPath := strings.Replace(f, filepath.Clean(sp), "", 1)
			tf := template.Must(template.New(filepath.Base(f)).ParseFiles(f))
			tfs = append(tfs, NewTemplate(relPath, tf))
		}

	}
	r.Templates = tfs
}

//getBuildDir generates a unique name for a temporary directory as well as creating the directory itself
func getBuildDir(p string) (string, error) {
	dname := p + buildDir + "/" + uuid.NewV1().String()

	err := os.MkdirAll(dname, os.FileMode(fileMode))
	if err != nil {
		return "", err
	}
	return dname, nil
}

func (r *Repo) SetConfig(c config.Config) {
	r.Config = c
}

func (r *Repo) Execute() error {
	for _, t := range r.Templates {
		err := t.Write(r.BuildDir, r.Config)
		if err != nil {
			return err
		}
	}
	return nil
}
