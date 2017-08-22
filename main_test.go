package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var (
	cliPath string
)

var _ = Describe("concourse-up", func() {
	BeforeSuite(func() {
		compilationVars := map[string]string{}

		file, err := os.Open("compilation-vars.json")
		Expect(err).To(Succeed())
		defer file.Close()

		err = json.NewDecoder(file).Decode(&compilationVars)
		Expect(err).To(Succeed())

		ldflags := []string{
			fmt.Sprintf("-X main.ConcourseUpVersion=%s", "0.0.0"),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseStemcellURL=%s", compilationVars["concourse_stemcell_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseStemcellVersion=%s", compilationVars["concourse_stemcell_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseStemcellSHA1=%s", compilationVars["concourse_stemcell_sha1"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseReleaseURL=%s", compilationVars["concourse_release_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseReleaseVersion=%s", compilationVars["concourse_release_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.ConcourseReleaseSHA1=%s", compilationVars["concourse_release_sha1"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.GardenReleaseURL=%s", compilationVars["garden_release_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.GardenReleaseVersion=%s", compilationVars["garden_release_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.GardenReleaseSHA1=%s", compilationVars["garden_release_sha1"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorStemcellURL=%s", compilationVars["director_stemcell_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorStemcellSHA1=%s", compilationVars["director_stemcell_sha1"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorStemcellVersion=%s", compilationVars["director_stemcell_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorCPIReleaseURL=%s", compilationVars["director_bosh_cpi_release_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorCPIReleaseVersion=%s", compilationVars["director_bosh_cpi_release_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorCPIReleaseSHA1=%s", compilationVars["director_bosh_cpi_release_sha1"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorReleaseURL=%s", compilationVars["director_bosh_release_url"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorReleaseVersion=%s", compilationVars["director_bosh_release_version"]),
			fmt.Sprintf("-X github.com/EngineerBetter/concourse-up/bosh.DirectorReleaseSHA1=%s", compilationVars["director_bosh_release_sha1"]),
		}

		cliPath, err = gexec.Build("github.com/EngineerBetter/concourse-up", "-ldflags", strings.Join(ldflags, " "))
		Expect(err).ToNot(HaveOccurred(), "Error building source")
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("displays usage instructions on --help", func() {
		command := exec.Command(cliPath, "--help")
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred(), "Error running CLI: "+cliPath)
		Eventually(session).Should(gexec.Exit(0))
		Expect(session.Out).To(Say("Concourse-Up - A CLI tool to deploy Concourse CI"))
		Expect(session.Out).To(Say("deploy, d   Deploys or updates a Concourse"))
		Expect(session.Out).To(Say("destroy, x  Destroys a Concourse"))
	})
})
