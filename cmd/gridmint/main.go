package main

import (
	"os"
	"path/filepath"

	"github.com/gridironOne/gridmint/cmd/gridmint/commands"
	"github.com/gridironOne/gridmint/config"
	"github.com/tendermint/tendermint/cmd/tendermint/commands/debug"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(
		commands.InitFilesCmd,
		commands.ShowSequencer,
		commands.ShowNodeIDCmd,
		debug.DebugCmd,
		cli.NewCompletionCmd(rootCmd, true),
	)

	// Create & start node
	rootCmd.AddCommand(commands.NewRunNodeCmd())

	cmd := cli.PrepareBaseCmd(rootCmd, "DM", os.ExpandEnv(filepath.Join("$HOME", config.DefaultDymintDir)))
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
