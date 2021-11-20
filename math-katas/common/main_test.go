package common

import (
	"reflect"
	"testing"
)

var (
	kPrimesData = []struct {
		k, start, nd int
		n            []int
	}{
		{2, 14, 18, []int{14, 15, 18}},
	}

	primeFactorsData = []struct {
		x  int
		n  []int
		ok bool
	}{
		{14, []int{2, 7}, true},
		{15, []int{3, 5}, true},
		{15, []int{2, 5}, false},
	}
)

func TestKPrimes(t *testing.T) {
	for _, table := range kPrimesData {
		expected := table.n
		actual := CountKPrimes(table.k, table.start, table.nd)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Test of (%d,%d,%d) was incorrect, got: %v, want: %v.", table.k, table.start, table.nd, actual, expected)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, table := range primeFactorsData {
		expected := table.n
		actual := PrimeFactors(table.x)
		if reflect.DeepEqual(actual, expected) != table.ok {
			t.Errorf("Test of (%d) was incorrect, got: %v, want: %v.", table.n, actual, expected)
		}
	}
}
