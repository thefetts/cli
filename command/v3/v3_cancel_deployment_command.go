package v3

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccversion"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v3/shared"
)

//go:generate counterfeiter . V3CancelDeploymentActor

type V3CancelDeploymentActor interface {
	CancelDeploymentByAppNameAndSpace(appName string, spaceGUID string) (v3action.Warnings, error)
	CloudControllerAPIVersion() string
}

type CancelDeploymentCommand struct {
	RequiredArgs flag.AppName
	UI           command.UI
	Config       command.Config
	Actor        V3CancelDeploymentActor
	SharedActor  command.SharedActor
}

func (cmd *CancelDeploymentCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config

	client, uaaClient, _ := shared.NewClients(config, ui, true, "")
	// if err != nil {
	// 	if v3Err, ok := err.(ccerror.V3UnexpectedResponseError); ok && v3Err.ResponseCode == http.StatusNotFound {
	// 		return translatableerror.MinimumCFAPIVersionNotMetError{MinimumVersion: ccversion.MinVersionApplicationFlowV3}
	// 	}
	//
	// 	return err
	// }

	cmd.Actor = v3action.NewActor(client, config, sharedaction.NewActor(config), uaaClient)

	return nil
}

func (cmd CancelDeploymentCommand) Execute(args []string) error {
	cmd.UI.DisplayWarning(command.ExperimentalWarning)

	err := command.MinimumCCAPIVersionCheck(cmd.Actor.CloudControllerAPIVersion(), ccversion.MinVersionApplicationFlowV3)
	if err != nil {
		return err
	}
	err = cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	warnings, err := cmd.Actor.CancelDeploymentByAppNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().Name)
	cmd.UI.DisplayWarnings(warnings)

	return err
}
