package src

import (
	"testing"
)

func TestDomainResultToString(t *testing.T) {
	domainResult := DomainResult{
		Domain: "example.com",
		Files: []FileResult{
			{
				Filename: "file1",
				Found:    true,
			},
			{
				Filename: "file2",
				Found:    false,
			},
		},
	}
	expected := "example.com\n\tfile1: Found\n\tfile2: Not found\n"
	if domainResult.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, domainResult.ToString())
	}
}
