package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfManifestUpdater(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfManifestUpdater Suite")
}
