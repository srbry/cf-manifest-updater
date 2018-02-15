package manifest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/srbry/cf-manifest-updater/manifest"
)

var _ = Describe("#Update", func() {
	It("updates a cf manifest", func() {
		Ω(manifest.Update(oldManifest)).Should(Equal(newManifest))
		Ω(manifest.Update(oldManifest1)).Should(Equal(newManifest1))
		Ω(manifest.Update(oldManifest2)).Should(Equal(newManifest2))
		Ω(manifest.Update(oldManifest3)).Should(Equal(newManifest3))
		Ω(manifest.Update(oldManifest4)).Should(Equal(newManifest4))
		Ω(manifest.Update(oldManifest5)).Should(Equal(newManifest5))
		Ω(manifest.Update(oldManifest6)).Should(Equal(newManifest6))
		Ω(manifest.Update(oldManifest7)).Should(Equal(newManifest7))
	})
})
