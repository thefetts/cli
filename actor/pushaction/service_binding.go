package pushaction

import "sort"

func (actor Actor) BindServices(config ApplicationConfig) (ApplicationConfig, bool, Warnings, error) {
	var allWarnings Warnings
	var boundService bool
	appGUID := config.DesiredApplication.GUID
	for serviceInstanceName, serviceInstance := range config.DesiredServices {
		if _, ok := config.CurrentServices[serviceInstanceName]; !ok {

			//TODO: Sort based on position
			var someSlice []int
			someSlice = sort.Sort(serviceInstance.Position)

			warnings, err := actor.V2Actor.BindServiceByApplicationAndServiceInstance(appGUID, serviceInstance.PushServiceInstance.GUID)
			allWarnings = append(allWarnings, warnings...)
			if err != nil {
				return config, false, allWarnings, err
			}
			boundService = true
		}
	}

	config.CurrentServices = config.DesiredServices
	return config, boundService, allWarnings, nil
}
