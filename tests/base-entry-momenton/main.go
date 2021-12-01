package main

type minSec struct {
	minutes int
	seconds int
}

func convertToMinutesAndSeconds(s int) minSec {
	mod := s % 60
	return minSec{(s - mod) / 60, mod}
}
