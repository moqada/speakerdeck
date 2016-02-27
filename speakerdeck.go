// Package speakerdeck is Fetch info and Download Slide from Speaker Deck for Go.
package speakerdeck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	base       = "https://speakerdeck.com"
	reSlideURL = regexp.MustCompile(`^` + base + `/([\w_-]+)/([\w_-]+)`)
)

// User is spearkerdeck user info
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
}

// Category is speakerdeck category info
type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

// Slide is speakerdeck slide info
type Slide struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DownloadURL string    `json:"downloadUrl"`
	User        User      `json:"user"`
	Stars       int       `json:"stars"`
	Category    Category  `json:"category"`
	Published   time.Time `json:"published"`
	URL         string    `json:"url"`
}

// NewUser is initialize user
func NewUser(username, displayName string) User {
	user := User{Username: username, DisplayName: displayName}
	user.URL = base + "/" + username
	return user
}

// NewCategory is initialize category
func NewCategory(name, slug string) Category {
	cat := Category{Name: name, Slug: slug}
	cat.URL = base + "/c/" + slug
	return cat
}

// DownloadPDF downloads PDF of target slide
func (s *Slide) DownloadPDF() ([]byte, error) {
	res, err := http.Get(s.DownloadURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetSlide returns slide info from Speakerdeck slide page
func GetSlide(url string) (*Slide, error) {
	slug, err := parseSlideURL(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	details := doc.Find("#talk-details")

	sidebar := doc.Find(".sidebar")
	a := sidebar.Find("h2 a")
	user := NewUser(a.AttrOr("href", "")[1:], a.Text())

	cat := sidebar.Find(".category a")
	category := NewCategory(cat.Text(), strings.Split(cat.AttrOr("href", ""), "/")[2])

	stars, err := strconv.Atoi(strings.Split(sidebar.Find(".stargazers").Text(), " ")[0])
	if err != nil {
		return nil, err
	}
	published, err := time.Parse("January  2, 2006", details.Find("header mark").First().Text())
	if err != nil {
		return nil, err
	}

	return &Slide{
		Slug:        slug,
		Title:       details.Find("h1").Text(),
		Description: strings.TrimSpace(details.Find(".description").Text()),
		Published:   published,
		DownloadURL: doc.Find("#share_pdf").AttrOr("href", ""),
		User:        user,
		Category:    category,
		Stars:       stars,
		URL:         user.URL + "/" + slug,
	}, nil
}

// parseSlideURL returns slide's slug from url
func parseSlideURL(url string) (string, error) {
	matches := reSlideURL.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", fmt.Errorf("Invalid URL")
	}
	return matches[2], nil
}
