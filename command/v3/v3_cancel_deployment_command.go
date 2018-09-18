package v3

import (
	"errors"
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

type V3CancelDeploymentCommand struct {
	RequiredArgs flag.AppName `positional-args:"yes"`
	UI           command.UI
	Config       command.Config
	CancelDeploymentActor        V3CancelDeploymentActor
	SharedActor  command.SharedActor	
    usage        interface{} `usage:"CF_NAME v3-cancel-zdt-push APP_NAME"`
}

func (cmd *V3CancelDeploymentCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)

	ccClient, uaaClient, err := shared.NewClients(config, ui, true, "")
	if err != nil {
		//if v3Err, ok := err.(ccerror.V3UnexpectedResponseError); ok && v3Err.ResponseCode == http.StatusNotFound {
		//	return translatableerror.MinimumCFAPIVersionNotMetError{MinimumVersion: ccversion.MinVersionApplicationFlowV3}
		//}

		return err
	}

	cmd.CancelDeploymentActor = v3action.NewActor(ccClient, config, sharedActor, uaaClient)
	cmd.SharedActor = sharedActor

	return nil
}

func (cmd V3CancelDeploymentCommand) Execute(args []string) error {
	cmd.UI.DisplayWarning(command.ExperimentalWarning)

	err := cmd.validateArgs()
	if err != nil {
		return err
	}

	err = command.MinimumCCAPIVersionCheck(cmd.CancelDeploymentActor.CloudControllerAPIVersion(), ccversion.MinVersionApplicationFlowV3)
	if err != nil {
		return err
	}
	err = cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	_, err = cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	warnings, err := cmd.CancelDeploymentActor.CancelDeploymentByAppNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)
	return err
}

func (cmd V3CancelDeploymentCommand) validateArgs() error {
	if cmd.RequiredArgs.AppName == "" {
		return errors.New("No app name given")
	}
	return nil
}
