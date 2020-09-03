package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
//Toma un slice como argumento, y devuelve un puntero de tipo IssuesSearchResult
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	//Crea un string uniendo todos los slices separados por espacios, y luego hace el escape para dejarlo compatible con html
	q := url.QueryEscape(strings.Join(terms, " "))
	//Hace la llamada
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	//Crea una variable de tipo IssuesSearchResult
	var result IssuesSearchResult
	//Convierte el json al tipo IssuesSearchResult y lo guarda en result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	//Devuelbe el puntero a result
	return &result, nil
}
