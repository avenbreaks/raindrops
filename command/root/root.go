package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/avenbreaks/raindrops/command/backup"
	"github.com/avenbreaks/raindrops/command/bridge"
	"github.com/avenbreaks/raindrops/command/genesis"
	"github.com/avenbreaks/raindrops/command/helper"
	"github.com/avenbreaks/raindrops/command/ibft"
	"github.com/avenbreaks/raindrops/command/license"
	"github.com/avenbreaks/raindrops/command/monitor"
	"github.com/avenbreaks/raindrops/command/peers"
	"github.com/avenbreaks/raindrops/command/polybft"
	"github.com/avenbreaks/raindrops/command/polybftsecrets"
	"github.com/avenbreaks/raindrops/command/regenesis"
	"github.com/avenbreaks/raindrops/command/rootchain"
	"github.com/avenbreaks/raindrops/command/secrets"
	"github.com/avenbreaks/raindrops/command/server"
	"github.com/avenbreaks/raindrops/command/status"
	"github.com/avenbreaks/raindrops/command/txpool"
	"github.com/avenbreaks/raindrops/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
