package version_test

import (
	"os"
	"os/exec"

	"code.cloudfoundry.org/cli/version"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {
	Describe("VersionString", func() {
		When("passed no ldflags", func() {
			It("returns the default version", func() {
				Expect(version.VersionString()).To(Equal("0.0.0-unknown-version"))
			})
		})

		When("a custom version is set", func() {
			It("returns the custom version", func() {
				version.SetVersion("1.2.3")
				Expect(version.VersionString()).To(Equal("1.2.3"))
			})
		})
	})

	Describe("SetVersion", func() {
		It("sets the version for valid semver versions", func() {
			version.SetVersion("1.2.3")
			Expect(version.VersionString()).To(Equal("1.2.3"))
		})

		It("exits with status code 1 when given an invalid semver", func() {
			if os.Getenv("TEST_EXIT") == "1" {
				version.SetVersion("not-a-semver")
				Fail("Expected process to exit but it didn't")
			}

			cmd := exec.Command(os.Args[0], "-ginkgo.focus=exits with status code 1 when given an invalid semver")
			cmd.Env = append(os.Environ(), "TEST_EXIT=1")

			err := cmd.Run()
			Expect(err).To(HaveOccurred())

			exitErr, ok := err.(*exec.ExitError)
			Expect(ok).To(BeTrue())
			Expect(exitErr.ExitCode()).To(Equal(1))
		})
	})
})
