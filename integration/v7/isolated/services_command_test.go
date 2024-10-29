package isolated

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccversion"
	"code.cloudfoundry.org/cli/integration/assets/hydrabroker/config"
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/servicebrokerstub"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("services command", func() {
	var (
		userName string
	)

	BeforeEach(func() {
		userName, _ = helpers.GetCredentials()
	})

	Describe("help", func() {
		When("--help flag is set", func() {
			It("Displays command usage to output", func() {
				session := helpers.CF("services", "--help")

				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("services - List all service instances in the target space"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("cf services"))
				Eventually(session).Should(Say("ALIAS:"))
				Eventually(session).Should(Say("s"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("create-service, marketplace"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("has no services", func() {
		BeforeEach(func() {
			helpers.LoginCF()
			helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
		})

		It("tells the user 'no services found'", func() {
			session := helpers.CF("services")

			Eventually(session).Should(Say("Getting services in org %s / space %s as %s...", ReadOnlyOrg, ReadOnlySpace, userName))
			Eventually(session).Should(Say("No services found"))
			Eventually(session).Should(Exit(0))
		})
	})

	Context("has services and applications", func() {
		var (
			orgName   string
			spaceName string

			broker *servicebrokerstub.ServiceBrokerStub

			managedService1      string
			managedService2      string
			userProvidedService1 string
			userProvidedService2 string
			appName1             string
			appName2             string
		)

		BeforeEach(func() {
			orgName = helpers.NewOrgName()
			spaceName = helpers.NewSpaceName()
			helpers.SetupCF(orgName, spaceName)

			userProvidedService1 = helpers.PrefixedRandomName("UPS1")
			userProvidedService2 = helpers.PrefixedRandomName("UPS2")

			Eventually(helpers.CF("cups", userProvidedService1, "-p", `{"username": "admin", "password": "admin"}`)).Should(Exit(0))
			Eventually(helpers.CF("cups", userProvidedService2, "-p", `{"username": "admin", "password": "admin"}`)).Should(Exit(0))

			broker = servicebrokerstub.EnableServiceAccess()

			managedService1 = helpers.PrefixedRandomName("MANAGED1")
			managedService2 = helpers.PrefixedRandomName("MANAGED2")
			Eventually(helpers.CF("create-service", broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), managedService1)).Should(Exit(0))
			Eventually(helpers.CF("create-service", broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), managedService2)).Should(Exit(0))

			appName1 = helpers.PrefixedRandomName("APP1")
			appName2 = helpers.PrefixedRandomName("APP2")
			helpers.WithHelloWorldApp(func(appDir string) {
				Eventually(helpers.CF("push", appName1, "--no-start", "-p", appDir, "-b", "staticfile_buildpack", "--no-route")).Should(Exit(0))
				Eventually(helpers.CF("push", appName2, "--no-start", "-p", appDir, "-b", "staticfile_buildpack", "--no-route")).Should(Exit(0))
			})
			Eventually(helpers.CF("bind-service", appName1, managedService1)).Should(Exit(0))
			Eventually(helpers.CF("bind-service", appName1, managedService2)).Should(Exit(0))
			Eventually(helpers.CF("bind-service", appName1, userProvidedService1)).Should(Exit(0))
			Eventually(helpers.CF("bind-service", appName1, userProvidedService2)).Should(Exit(0))
			Eventually(helpers.CF("bind-service", appName2, managedService2)).Should(Exit(0))
			Eventually(helpers.CF("bind-service", appName2, userProvidedService2)).Should(Exit(0))
		})

		AfterEach(func() {
			Eventually(helpers.CF("unbind-service", appName1, managedService1)).Should(Exit(0))
			Eventually(helpers.CF("unbind-service", appName1, managedService2)).Should(Exit(0))
			Eventually(helpers.CF("unbind-service", appName1, userProvidedService1)).Should(Exit(0))
			Eventually(helpers.CF("unbind-service", appName1, userProvidedService2)).Should(Exit(0))
			Eventually(helpers.CF("unbind-service", appName2, managedService2)).Should(Exit(0))
			Eventually(helpers.CF("unbind-service", appName2, userProvidedService2)).Should(Exit(0))
			Eventually(helpers.CF("delete-service", managedService1, "-f")).Should(Exit(0))
			Eventually(helpers.CF("delete-service", managedService2, "-f")).Should(Exit(0))
			broker.Forget()
			helpers.QuickDeleteOrg(orgName)
		})

		When("CAPI version is < 2.125.0", func() {
			It("displays all service information", func() {
				session := helpers.CF("services")
				Eventually(session).Should(Say("Getting services in org %s / space %s as %s...", orgName, spaceName, userName))
				Eventually(session).Should(Say(`name\s+service\s+plan\s+bound apps\s+last operation\s+broker`))
				Eventually(session).Should(Say(`%s\s+%s\s+%s\s+%s\s+%s\s`, managedService1, broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), appName1, "create succeeded"))
				Eventually(session).Should(Say(`%s\s+%s\s+%s\s+%s, %s\s+%s\s`, managedService2, broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), appName1, appName2, "create succeeded"))
				Eventually(session).Should(Say(`%s\s+%s\s+%s`, userProvidedService1, "user-provided", appName1))
				Eventually(session).Should(Say(`%s\s+%s\s+%s, %s`, userProvidedService2, "user-provided", appName1, appName2))
				Eventually(session).Should(Exit(0))
			})
		})

		When("CAPI version is >= 2.125.0 (broker name is available in summary endpoint)", func() {
			It("displays all service information", func() {
				session := helpers.CF("services")
				Eventually(session).Should(Say("Getting services in org %s / space %s as %s...", orgName, spaceName, userName))
				Eventually(session).Should(Say(`name\s+service\s+plan\s+bound apps\s+last operation\s+broker`))
				Eventually(session).Should(Say(`%s\s+%s\s+%s\s+%s\s+%s\s+%s`, managedService1, broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), appName1, "create succeeded", broker.Name))
				Eventually(session).Should(Say(`%s\s+%s\s+%s\s+%s, %s\s+%s\s+%s`, managedService2, broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), appName1, appName2, "create succeeded", broker.Name))
				Eventually(session).Should(Say(`%s\s+%s\s+%s`, userProvidedService1, "user-provided", appName1))
				Eventually(session).Should(Say(`%s\s+%s\s+%s, %s`, userProvidedService2, "user-provided", appName1, appName2))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("has shared service instances", func() {
		var (
			managedService, appNameOnSpaceA, appNameOnSpaceB, orgName string
		)

		BeforeEach(func() {
			orgName = helpers.NewOrgName()
			spaceA := helpers.NewSpaceName()
			spaceB := helpers.NewSpaceName()
			managedService = helpers.PrefixedRandomName("MANAGED1")
			appNameOnSpaceA = helpers.PrefixedRandomName("APP1")
			appNameOnSpaceB = helpers.PrefixedRandomName("APP1")

			helpers.SetupCF(orgName, spaceA)
			helpers.CreateOrgAndSpace(orgName, spaceB)
			broker := servicebrokerstub.New().WithPlans(2).EnableServiceAccess()
			helpers.CreateManagedServiceInstance(broker.FirstServiceOfferingName(), broker.FirstServicePlanName(), managedService)

			helpers.WithHelloWorldApp(func(appDir string) {
				Eventually(helpers.CF("push", appNameOnSpaceA, "--no-start", "-p", appDir, "-b", "staticfile_buildpack", "--no-route")).Should(Exit(0))
			})
			Eventually(helpers.CF("bind-service", appNameOnSpaceA, managedService)).Should(Exit(0))
			Eventually(helpers.CF("share-service", managedService, "-s", spaceB)).Should(Exit(0))

			helpers.TargetOrgAndSpace(orgName, spaceB)
			helpers.WithHelloWorldApp(func(appDir string) {
				Eventually(helpers.CF("push", appNameOnSpaceB, "--no-start", "-p", appDir, "-b", "staticfile_buildpack", "--no-route")).Should(Exit(0))
			})
			Eventually(helpers.CF("bind-service", appNameOnSpaceB, managedService)).Should(Exit(0))
			helpers.TargetOrgAndSpace(orgName, spaceA)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		It("should not output bound apps in the shared spaces", func() {
			session := helpers.CF(command)
			Eventually(session).Should(Exit(0))
			Expect(session).To(SatisfyAll(
				Say(managedService),
				Say(appNameOnSpaceA),
				Not(Say(appNameOnSpaceB)),
			))
		})
	})
})
