package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSnafoozSolver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SnafoozSolver Suite")
}
