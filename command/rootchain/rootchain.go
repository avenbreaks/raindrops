package rootchain

import (
	"github.com/spf13/cobra"

	"github.com/avenbreaks/raindrops/command/rootchain/deploy"
	"github.com/avenbreaks/raindrops/command/rootchain/fund"
	"github.com/avenbreaks/raindrops/command/rootchain/server"
)

// GetCommand creates "rootchain" helper command
func GetCommand() *cobra.Command {
	rootchainCmd := &cobra.Command{
		Use:   "rootchain",
		Short: "Top level rootchain helper command.",
	}

	rootchainCmd.AddCommand(
		// rootchain server
		server.GetCommand(),
		// rootchain deploy
		deploy.GetCommand(),
		// rootchain fund
		fund.GetCommand(),
	)

	return rootchainCmd
}
