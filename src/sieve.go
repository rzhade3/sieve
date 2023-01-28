package src

import (
	"bytes"
	"net/url"
	"text/template"
)

type DomainResult struct {
	Domain string
	Files  []FileResult
}

type FileResult struct {
	Filename string
	Found    bool
}

// Checks a domain to see if a list of files is exposed
func CheckDomain(domain string, files *[]File) []FileResult {
	var results []FileResult
	for _, file := range *files {
		file_url, err := url.JoinPath("https://", domain, file.Filename)
		if err != nil {
			results = append(results, FileResult{Filename: file.Filename, Found: false})
			continue
		}
		content, err := FetchUrl(file_url)
		if err != nil {
			results = append(results, FileResult{Filename: file.Filename, Found: false})
			continue
		}
		results = append(results, FileResult{Filename: file.Filename, Found: ValidateContent(content, file.Filetype)})
	}
	return results
}

// Stringifies a DomainResult
func (d DomainResult) ToString() string {
	var result bytes.Buffer
	tmpl, _ := template.New("domainresult").Parse("{{.Domain}}\n{{range .Files}}\t{{.Filename}}: {{if .Found}}Found{{else}}Not found{{end}}\n{{end}}")
	tmpl.Execute(&result, d)
	return result.String()
}
