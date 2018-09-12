package v2_test

import (
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	. "code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = FDescribe("CreateSpaceCommand", func() {
	var (
		fakeConfig      *commandfakes.FakeConfig
		fakeActor       *v2fakes.FakeCreateSpaceActor
		fakeSharedActor *commandfakes.FakeSharedActor
		testUI          *ui.UI
		spaceName       string
		cmd             CreateSpaceCommand

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v2fakes.FakeCreateSpaceActor)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		spaceName = "some-space"

		cmd = CreateSpaceCommand{
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			SharedActor:  fakeSharedActor,
			RequiredArgs: flag.Space{Space: spaceName},
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	//TODO: Delete me when this command is no longer experimental
	When("The experimental flag is not set", func() {
		It("is experimental", func() {
			Expect(executeErr).To(MatchError(translatableerror.UnrefactoredCommandError{}))
		})
	})

	When("the experimental flag is set", func() {
		var checkTargetedOrg bool

		BeforeEach(func() {
			fakeConfig.ExperimentalReturns(true)
		})

		It("checks for user being logged in", func() {
			var checkTargetedSpace bool
			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			checkTargetedOrg, checkTargetedSpace = fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedSpace).To(BeFalse())
		})

		When("user is not logged in", func() {
			expectedErr := errors.New("please make a require-current-user function from checktarget and currentuser")

			BeforeEach(func() {
				fakeSharedActor.CheckTargetReturns(expectedErr)
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError(expectedErr))
			})
		})

		When("user is logged in", func() {
			var username string

			BeforeEach(func() {
				username = "some-guy"

				fakeSharedActor.CheckTargetReturns(nil)

				fakeConfig.CurrentUserReturns(configv3.User{
					Name: username,
				}, nil)
			})

			When("user specifies an org using the -o flag", func() {
				var specifiedOrgName string

				BeforeEach(func() {
					specifiedOrgName = "specified-org"
					cmd.Organization = specifiedOrgName

					fakeConfig.HasTargetedOrganizationReturns(true)
					fakeConfig.TargetedOrganizationReturns(configv3.Organization{
						GUID: "irrelevant-guid",
						Name: "irrelevant-name",
					})
				})

				It("does not require a targeted org", func() {
					Expect(checkTargetedOrg).To(BeFalse())
				})

				It("uses the specified org, not the targeted org", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(fakeConfig.CurrentUserCallCount()).To(Equal(1))
					Expect(testUI.Out).To(Say(`Creating space %s in org %s as %s\.\.\.`, spaceName, specifiedOrgName, username))

					Expect(testUI.Out).To(Say("OK\n\n"))

					Expect(fakeActor.CreateSpaceCallCount()).To(Equal(1))
					inputSpace, inputOrg, _ := fakeActor.CreateSpaceArgsForCall(0)
					Expect(inputSpace).To(Equal(spaceName))
					Expect(inputOrg).To(Equal(specifiedOrgName))
				})
			})

			When("no org is specified using the -o flag", func() {
				It("requires a targeted org", func() {
					Expect(checkTargetedOrg).To(BeTrue())
				})

				When("checking target fails", func() {
					BeforeEach(func() {
						fakeSharedActor.CheckTargetReturns(errors.New("check target error"))
					})

					It("returns an error", func() {
						Expect(executeErr).To(MatchError("check target error"))
					})
				})

				When("no org is targeted", func() {
					BeforeEach(func() {
						fakeConfig.HasTargetedOrganizationReturns(false)
					})

					It("returns an error", func() {

					})
				})

				When("an org is targeted", func() {
					var targetedOrgName string

					BeforeEach(func() {
						fakeConfig.HasTargetedOrganizationReturns(true)
						fakeConfig.TargetedOrganizationReturns(configv3.Organization{
							GUID: "some-org-guid",
							Name: targetedOrgName,
						})
					})

					When("creating the space succeeds", func() {
						BeforeEach(func() {
							fakeActor.CreateSpaceReturns(
								v2action.Space{
									GUID: "fake-space-id",
								},
								v2action.Warnings{"warn-1", "warn-2"},
								nil,
							)
						})

						It("creates the org and displays warnings", func() {
							Expect(executeErr).ToNot(HaveOccurred())
							Expect(fakeConfig.CurrentUserCallCount()).To(Equal(1))
							Expect(testUI.Out).To(Say(`Creating space %s in org %s as %s\.\.\.`, spaceName, targetedOrgName, username))

							Expect(testUI.Err).To(Say("warn-1\nwarn-2\n"))
							Expect(testUI.Out).To(Say("OK\n\n"))

							Expect(fakeActor.CreateSpaceCallCount()).To(Equal(1))
							inputSpace, inputOrg, quota := fakeActor.CreateSpaceArgsForCall(0)
							Expect(inputSpace).To(Equal(spaceName))
							Expect(inputOrg).To(Equal(targetedOrgName))
							Expect(quota).To(BeEmpty())
						})

					})
				})

				When("quota is not specified", func() {

				})

				When("quota is specified", func() {

				})
			})
		})

	})

})
