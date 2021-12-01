package main

import (
	"strings"
	"time"
)

const (
	inputLayout  = "2 January 2006"
	outputLayout = "2006-01-02"
)

func formatDate(s string) string {
	ordinal := []string{"nd", "rd", "st", "th"}
	form := inputLayout
	for _, o := range ordinal {
		if strings.Index(s, o+" ") > -1 {
			s = strings.Replace(s, o+" ", " ", -1)
			break
		}
	}
	t, _ := time.Parse(form, s)
	return t.Format(outputLayout)
}
