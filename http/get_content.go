package http

import (
	"fmt"
	"io"
	"net/http"
)

func GetContent(url, userAgent string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if (resp.StatusCode < 200) || (resp.StatusCode >= 300) {
		return "", fmt.Errorf("status code %d does not indicate success", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
