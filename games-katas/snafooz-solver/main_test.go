package main

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func tst(pieces [6][6][6]int) bool {
	solution := SolveSnafooz(pieces)
	if len(solution) != 6 {
		print("Expected length to be 6 but was", len(solution))
		return false
	}
	return reflect.DeepEqual(solution, pieces)
}

var _ = Describe("snafooz cube puzzle", func() {
	It("example puzzle1", func() {
		pieces := [6][6][6]int{{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 1, 1}},

			{{0, 1, 0, 0, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 0, 1}},

			{{0, 0, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 0, 0}},

			{{0, 0, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 0, 0, 1, 0}},

			{{0, 0, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 1, 1}},

			{{0, 0, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 0, 0, 1, 0}}}

		Expect(tst(pieces)).To(Equal(true))
	})

	It("example puzzle2", func() {
		pieces := [6][6][6]int{{{0, 0, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 0, 0}},

			{{1, 1, 0, 0, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1, 1}},

			{{0, 0, 1, 0, 0, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 0, 0, 1, 0}},

			{{0, 0, 1, 1, 0, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 0, 0}},

			{{0, 0, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 0, 0, 1, 1}},

			{{0, 0, 1, 1, 0, 1},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 0, 0, 1, 1}}}

		Expect(tst(pieces)).To(Equal(true))
	})

	It("example puzzle3", func() {
		pieces := [6][6][6]int{{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 1, 0}},

			{{0, 1, 0, 1, 0, 0},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 1, 0}},

			{{1, 1, 0, 0, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{0, 0, 1, 1, 0, 0}},

			{{0, 0, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{0, 1, 0, 0, 1, 1}},

			{{0, 0, 1, 1, 0, 0},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 0, 0}},

			{{1, 1, 0, 0, 1, 1},
				{1, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 0, 1}}}

		Expect(tst(pieces)).To(Equal(true))
	})
})
