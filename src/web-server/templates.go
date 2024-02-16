package web_server

import (
	"fmt"
	"html/template"
)

type templatePathMap struct {
	IndexPage    string
	Expanses     string
	ExpanseTypes string
	Auth         string
}

var tpm templatePathMap = setupTemplatePathMap()

func setupTemplatePathMap() templatePathMap {
	templatesPath := getTemplateFolderPath()

	tpm := templatePathMap{}
	tpm.IndexPage = fmt.Sprintf("%s/pages/index.go.html", templatesPath)
	tpm.Expanses = fmt.Sprintf("%s/parts/expenses.go.html", templatesPath)

	tpm.ExpanseTypes = fmt.Sprintf("%s/parts/expense-types.go.html", templatesPath)

	tpm.Auth = fmt.Sprintf("%s/parts/auth.go.html", templatesPath)

	return tpm
}

func getIndexPageTemplate() (*template.Template, error) {
	return template.ParseFiles(tpm.IndexPage, tpm.Expanses, tpm.ExpanseTypes, tpm.Auth)
}
