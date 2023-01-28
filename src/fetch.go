package src

import (
	"crypto/tls"
	"io"
	"net/http"
)

func FetchUrl(url string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "sieve/0.1")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(body), err
}
