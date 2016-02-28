package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/LK4D4/dolstats"
)

type sortedCases []dolstats.Case

func (c sortedCases) Len() int {
	return len(c)
}

func (c sortedCases) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c sortedCases) Less(i, j int) bool {
	return strings.Compare(c[i].Number, c[j].Number) == -1
}
func main() {
	head := 10
	if len(os.Args) > 1 {
		h, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Usage: %s <top number>", os.Args[0])
		}
		head = h
	}
	y, m, d := time.Now().Date()
	d--
	f := dolstats.Filter{
		From: time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC),
		To:   time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC),
	}
	cases, err := dolstats.GetCases(f)
	if err != nil {
		log.Fatalf("error getting cases: %v", err)
	}
	sort.Sort(sort.Reverse(sortedCases(cases)))
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Number\tPosted Date\tApproval Date\tProcessing Time\tOccupation\tCompany")
	if len(cases) == 0 {
		os.Exit(0)
	}
	for _, c := range cases[:head] {
		var yp, mp, dp, yc, mc, dc int
		if _, err := fmt.Sscanf(c.PostedDate, "%d-%d-%d", &mp, &dp, &yp); err != nil {
			log.Fatalf("error scanning posted date: %v", err)
		}
		if _, err := fmt.Sscanf(c.ApprovalDate, "%d-%d-%d", &mc, &dc, &yc); err != nil {
			log.Fatalf("error scanning certified date: %v", err)
		}
		postedDate := time.Date(yp, time.Month(mp), dp, 0, 0, 0, 0, time.UTC)
		certDate := time.Date(yc, time.Month(mc), dc, 0, 0, 0, 0, time.UTC)
		processingTime := certDate.Sub(postedDate)
		days := processingTime.Hours() / 24
		fmt.Fprintf(w, "%s\t%s\t%s\t%.0f days ~ %.0f months\t%s\t%s\n", c.Number, c.PostedDate, c.ApprovalDate, days, days/30, c.Job, c.Employer)
	}
	w.Flush()
}
