package main

import (
	"reflect"
	"testing"
)

var (
	data = []struct {
		seconds int
		ms      minSec
		ok      bool
	}{
		{62, minSec{1, 2}, true},
		{2, minSec{0, 2}, true},
		{60, minSec{1, 0}, true},
		{122, minSec{2, 2}, true},
		{183, minSec{3, 3}, true},
		{3607, minSec{60, 7}, true},
	}
)

func TestConvertToMinutesAndSeconds(t *testing.T) {
	for _, table := range data {
		expected := table.ms
		actual := convertToMinutesAndSeconds(table.seconds)
		if reflect.DeepEqual(actual, expected) != table.ok {
			t.Errorf("Test of (%v,%d) was incorrect, got: %v, want: %v.", table.seconds, table.ms, actual, expected)
		}
	}
}
