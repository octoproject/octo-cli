package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	git "github.com/go-git/go-git/v5"
	"github.com/octoproject/octo-cli/config"
)

const (
	octoTemplatesDir              = "templates"
	tmpDir                        = "/tmp/templates"
	defaultOctoTemplateRepository = "https://github.com/octoproject/templates"
)

var (
	ErrUnknownServiceType = errors.New("service type unknown")
)

// New generate new function
func NewFunction(s *config.Service) error {
	serviceDir := s.ServiceName
	// fetch octo templates if not exist
	err := fetchTemplates()
	if err != nil {
		return err
	}

	// copy function template to service folder
	err = copy(filepath.Join(octoTemplatesDir, "node12"), serviceDir)
	if err != nil {
		return err
	}

	// copy package-lock.json to service folder
	err = copy(filepath.Join(octoTemplatesDir, "package.json"),
		filepath.Join(s.ServiceName, "function", "package.json"))
	if err != nil {
		return err
	}

	// copy package-lock.json to service folder
	err = copy(filepath.Join(octoTemplatesDir, "package-lock.json"),
		filepath.Join(s.ServiceName, "function", "package-lock.json"))
	if err != nil {
		return err
	}

	// generate function handler
	err = generateHandler(s)
	if err != nil {
		return err
	}
	fmt.Printf("function: %s is created.\n", s.ServiceName)

	// move handler.js to service folder
	err = os.Rename("handler.js", filepath.Join(s.ServiceName, "function", "handler.js"))
	if err != nil {
	}
	return nil
}

// generateHandler will create new handler.js based on template file
func generateHandler(s *config.Service) error {
	var t *template.Template
	var err error

	switch s.ServiceType {
	case "read":
		t, err = template.ParseFiles(filepath.Join(octoTemplatesDir, "read-template.tmpl"))
		if err != nil {
			return err
		}
	case "write":
		t, err = template.ParseFiles(filepath.Join(octoTemplatesDir, "write-template.tmpl"))
		if err != nil {
			return err
		}
	default:
		return ErrUnknownServiceType
	}

	//create handler file
	f, err := os.Create("handler.js")
	if err != nil {
		return err
	}

	err = t.Execute(f, s)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}

//fetchTemplates fetch template if not exist
func fetchTemplates() error {
	if !exists(octoTemplatesDir) {
		fmt.Println("template is not exist, fetching the template.")
		err := cloneDefaultTemplate()
		if err != nil {
			return err
		}

		err = os.Rename(tmpDir, octoTemplatesDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// clone default template repo
func cloneDefaultTemplate() error {
	_, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:      defaultOctoTemplateRepository,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}
