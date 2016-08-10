package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands/flags"
)

type DeleteSharedDomainCommand struct {
	RequiredArgs flags.Domain `positional-args:"yes"`
	Force        bool         `short:"f" description:"Force deletion without confirmation"`
	usage        interface{}  `usage:"CF_NAME delete-shared-domain DOMAIN [-f]"`
}

func (_ DeleteSharedDomainCommand) Setup() error {
	return nil
}

func (_ DeleteSharedDomainCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}