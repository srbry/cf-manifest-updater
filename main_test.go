package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var pathToCLI string

var _ = BeforeSuite(func() {
	var err error
	pathToCLI, err = gexec.Build("github.com/srbry/cf-manifest-updater")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = Describe("cf-manifest-updater", func() {
	var (
		err     error
		session *gexec.Session
	)

	BeforeEach(func() {
		command := exec.Command(pathToCLI, "fixtures/old_manifest.yml")
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	})

	AfterEach(func() {
		session.Kill()
	})

	It("show the updated manifest", func() {
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say(`---
applications:
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example
  routes:
  - route: test.example1.com
  - route: test.example2.com
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example
  routes:
  - route: example.example1.com
  - route: example.example2.com
`))
	})
})
