package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKPrimes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KPrimes Suite")
}
