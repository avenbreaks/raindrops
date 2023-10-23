package ibft

import (
	"github.com/avenbreaks/raindrops/command/helper"
	"github.com/avenbreaks/raindrops/command/ibft/candidates"
	"github.com/avenbreaks/raindrops/command/ibft/propose"
	"github.com/avenbreaks/raindrops/command/ibft/quorum"
	"github.com/avenbreaks/raindrops/command/ibft/snapshot"
	"github.com/avenbreaks/raindrops/command/ibft/status"
	_switch "github.com/avenbreaks/raindrops/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
