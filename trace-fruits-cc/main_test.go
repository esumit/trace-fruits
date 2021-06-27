package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestTraceEvents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trace Events Suite")
}

var _ = Describe(`TraceEvent`, func() {

})
