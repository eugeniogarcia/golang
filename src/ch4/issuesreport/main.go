// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"ch4/github"
)

//Crea el template. En el template tenemos un listado que especificamos con range
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//Funcion helper
func haceDias(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//Crea el template con Must, le llama issuelist, le aplica una lambda que reemplaza daysAgo por retrulado de una funcion
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": haceDias}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//!-exec

func noMust() {
	//!+parse
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": haceDias}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	//!-parse
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
