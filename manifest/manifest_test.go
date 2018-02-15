package manifest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/srbry/cf-manifest-updater/manifest"
)

var _ = Describe("#Update", func() {
	Context("when using manifest 0", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest)).Should(Equal(newManifest))
		})
	})

	Context("when using manifest 1", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest1)).Should(Equal(newManifest1))
		})
	})

	Context("when using manifest 2", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest2)).Should(Equal(newManifest2))
		})
	})

	Context("when using manifest 3", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest3)).Should(Equal(newManifest3))
		})
	})

	Context("when using manifest 4", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest4)).Should(Equal(newManifest4))
		})
	})

	Context("when using manifest 5", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest5)).Should(Equal(newManifest5))
		})
	})

	Context("when using manifest 6", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest6)).Should(Equal(newManifest6))
		})
	})

	Context("when using manifest 7", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest7)).Should(Equal(newManifest7))
		})
	})

	Context("when using manifest 8", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest8)).Should(Equal(newManifest8))
		})
	})

	Context("when using manifest 9", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest9)).Should(Equal(newManifest9))
		})
	})

	Context("when using manifest 10", func() {
		It("updates a cf manifest", func() {
			Ω(manifest.Update(oldManifest10)).Should(Equal(newManifest10))
		})
	})
})
