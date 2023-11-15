package e2e_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "../pkg/helpers"
	"os"
)

const baseDir = "/tmp/multi-git"

var repoList string

var _ = Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		Ω(err).Should(BeNil())
	}

	BeforeEach(func() {
		removeAll()
		err = CreateDir(baseDir, "", false)
		Ω(err).Should(BeNil())
	})

	AfterSuite(removeAll)
})