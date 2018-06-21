package pushaction_test

import (
	"errors"

	. "code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/pushaction/pushactionfakes"
	"code.cloudfoundry.org/cli/actor/v2action"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Binding Services", func() {
	var (
		actor       *Actor
		fakeV2Actor *pushactionfakes.FakeV2Actor
	)

	BeforeEach(func() {
		fakeV2Actor = new(pushactionfakes.FakeV2Actor)
		actor = NewActor(fakeV2Actor, nil, nil)
	})

	Describe("BindServices", func() {
		var (
			config ApplicationConfig

			returnedConfig ApplicationConfig
			boundServices  bool
			warnings       Warnings
			executeErr     error
		)

		BeforeEach(func() {
			config = ApplicationConfig{}
			config.DesiredApplication.GUID = "some-app-guid"

			service1 := v2action.ServiceInstance{
				GUID: "instance_1_guid",
			}
			service2 := v2action.ServiceInstance{
				GUID: "instance_2_guid",
			}
			service3 := v2action.ServiceInstance{
				GUID: "instance_3_guid",
			}

			config.CurrentServices = map[string]Service{
				"service_instance_1": {
					PushServiceInstance: service1,
					Position:            0,
				},
			}

			config.DesiredServices = map[string]Service{
				"service_instance_1": {
					PushServiceInstance: service1,
					Position:            0,
				},
				"service_instance_2": {
					PushServiceInstance: service2,
					Position:            0,
				},
				"service_instance_3": {
					PushServiceInstance: service3,
					Position:            0,
				},
			}
		})

		JustBeforeEach(func() {
			returnedConfig, boundServices, warnings, executeErr = actor.BindServices(config)
		})

		Context("when binding services is successful", func() {
			BeforeEach(func() {
				fakeV2Actor.BindServiceByApplicationAndServiceInstanceReturnsOnCall(0, v2action.Warnings{"service-instance-warning-1"}, nil)
				fakeV2Actor.BindServiceByApplicationAndServiceInstanceReturnsOnCall(1, v2action.Warnings{"service-instance-warning-2"}, nil)
			})

			It("it updates CurrentServices to match DesiredServices", func() {
				service1 := v2action.ServiceInstance{
					GUID: "instance_1_guid",
				}
				service2 := v2action.ServiceInstance{
					GUID: "instance_2_guid",
				}
				service3 := v2action.ServiceInstance{
					GUID: "instance_3_guid",
				}

				Expect(executeErr).ToNot(HaveOccurred())
				Expect(warnings).To(ConsistOf("service-instance-warning-1", "service-instance-warning-2"))
				Expect(boundServices).To(BeTrue())
				Expect(returnedConfig.CurrentServices).To(Equal(map[string]Service{
					"service_instance_1": {
						PushServiceInstance: service1,
					},
					"service_instance_2": {
						PushServiceInstance: service2,
					},
					"service_instance_3": {
						PushServiceInstance: service3,
					},
				}))

				var serviceInstanceGUIDs []string
				Expect(fakeV2Actor.BindServiceByApplicationAndServiceInstanceCallCount()).To(Equal(2))
				appGUID, serviceInstanceGUID := fakeV2Actor.BindServiceByApplicationAndServiceInstanceArgsForCall(0)
				Expect(appGUID).To(Equal("some-app-guid"))
				serviceInstanceGUIDs = append(serviceInstanceGUIDs, serviceInstanceGUID)

				appGUID, serviceInstanceGUID = fakeV2Actor.BindServiceByApplicationAndServiceInstanceArgsForCall(1)
				Expect(appGUID).To(Equal("some-app-guid"))
				serviceInstanceGUIDs = append(serviceInstanceGUIDs, serviceInstanceGUID)

				Expect(serviceInstanceGUIDs).To(ConsistOf("instance_2_guid", "instance_3_guid"))
			})
		})

		Context("when binding services fails", func() {
			BeforeEach(func() {
				fakeV2Actor.BindServiceByApplicationAndServiceInstanceReturns(v2action.Warnings{"service-instance-warning-1"}, errors.New("some-error"))
			})

			It("it returns the error", func() {
				Expect(executeErr).To(MatchError("some-error"))
				Expect(warnings).To(ConsistOf("service-instance-warning-1"))
				Expect(boundServices).To(BeFalse())
			})
		})
	})
})
