package manifest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/srbry/cf-manifest-updater/manifest"
)

const oldManifest = `name: example
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
`

const newManifest = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: example.example.com
`

var _ = Describe("#Update", func() {
	It("updates a cf manifest", func() {
		Ω(manifest.Update(oldManifest)).Should(Equal(newManifest))
	})
})
