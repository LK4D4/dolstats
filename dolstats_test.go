package dolstats

import (
	"testing"
	"time"
)

func TestFilterDateCompany(t *testing.T) {
	cs, err := GetCases(Filter{
		From:     time.Date(2016, 2, 25, 0, 0, 0, 0, time.UTC),
		To:       time.Date(2016, 2, 26, 0, 0, 0, 0, time.UTC),
		Employer: "GOOGLE",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 46 {
		t.Fatalf("expected 46 results, got %d", len(cs))
	}
}

func TestFilterDateState(t *testing.T) {
	cs, err := GetCases(Filter{
		From:  time.Date(2016, 2, 25, 0, 0, 0, 0, time.UTC),
		To:    time.Date(2016, 2, 26, 0, 0, 0, 0, time.UTC),
		State: "WI",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 27 {
		t.Fatalf("expected 27 results, got %d", len(cs))
	}
}

func TestFilterNumber(t *testing.T) {
	num := "A-15265-20518"
	job := "Poultry Processing Worker"
	cs, err := GetCases(Filter{
		Number: num,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 1 {
		t.Fatalf("expected 1 results, got %d", len(cs))
	}
	c := cs[0]
	if c.Number != num {
		t.Fatalf("expected case number %s, got %s", num, c.Number)
	}
	if c.Job != job {
		t.Fatalf("expected job to be %s, got %s", job, c.Job)
	}
}
