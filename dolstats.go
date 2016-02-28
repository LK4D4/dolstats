// Package dolstats is basic API to dolstats.com website - "database" of
// all processed PERM applications by Department of Labour. It doesn't require
// any registration.
package dolstats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const searchURL = "http://dolstats.com/searchAjax"

// Case represents PERM case which was processed by Department of Labour.
type Case struct {
	Number       string `json:"cn"`
	PostedDate   string `json:"cCD"`
	ApprovalDate string `json:"pD"`
	Job          string `json:"pT"`
	Employer     string `json:"fN"`
	State        string `json:"s"`
	Status       string `json:"cR"`
}

type answer struct {
	Result []Case
}

// Filter specifies search parameters.
type Filter struct {
	From     time.Time
	To       time.Time
	Number   string
	Employer string
	State    string
	Status   string
}

func getDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%d-%d", m, d, y)
}

func getURL(f Filter) string {
	v := &url.Values{}
	if !f.From.IsZero() {
		v.Set("from-date", getDate(f.From))
	}
	if !f.To.IsZero() {
		v.Set("to-date", getDate(f.To))
	}
	v.Set("case-types", "ALL")
	if f.Number != "" {
		v.Set("case-number", f.Number)
	}
	if f.Employer != "" {
		v.Set("employer-name", f.Employer)
	}
	if f.Status != "" {
		v.Set("status", f.Status)
	}
	if f.State != "" {
		v.Set("state", f.State)
	}
	return searchURL + "?" + v.Encode()
}

// GetCases returns all processed cases which satisfy f.
func GetCases(f Filter) ([]Case, error) {
	url := getURL(f)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ans := answer{}
	if err := json.NewDecoder(resp.Body).Decode(&ans); err != nil {
		return nil, err
	}
	return ans.Result, nil
}
