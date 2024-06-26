package scraper

import (
	"log/slog"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeLogin(url, username, password string) (valid bool, err error) {
	var loginToken string
	var actionURL string

	var c = colly.NewCollector()

	formData := make(map[string]string)

	c.OnHTML("form#login", func(e *colly.HTMLElement) {
		// Hidden value used to prevent outside access
		loginToken = e.ChildAttr("input[name=logintoken]", "value")

		actionURL = e.Attr("action")

		// Fill in the form fields.
		formData["username"] = username
		formData["password"] = password
		formData["logintoken"] = loginToken

	})

	c.OnError(func(r *colly.Response, e error) {
		slog.Error("Request error:", e)
		err = e
	})

	err = c.Visit(url)

	if err != nil {

		slog.Error("Error visiting URL:", err)
	}

	c.OnResponse(func(r *colly.Response) {
		responseURL := r.Request.URL
		slog.Info("Response received from", "body:", string(responseURL.Path))
		slog.Info("BOOL", "B:", responseURL.Path)
		if strings.Contains(responseURL.Path, "/my/") {
			valid = true
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		slog.Error("Request error:", e)
		err = e
	})

	err = c.Post(actionURL, formData)

	if err != nil {
		slog.Error("Error submitting form:", err)
	}

	return
}
