package goal

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xushuhui/goal/config"
	"github.com/xushuhui/goal/internal/command/create"
	"github.com/xushuhui/goal/internal/command/gen"
	"github.com/xushuhui/goal/internal/command/new"
	"github.com/xushuhui/goal/internal/command/run"
	"github.com/xushuhui/goal/internal/command/upgrade"
)

var CmdRoot = &cobra.Command{
	Use:     "goal",
	Example: "goal new demo-api",
	Short:   "Goal \n ",
	Version: fmt.Sprintf("Goal %s - Copyright (c) 2023 Goal\nReleased under the MIT License.\n\n", config.Version),
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
	CmdRoot.AddCommand(create.CmdCreate)
	CmdRoot.AddCommand(run.CmdRun)
	CmdRoot.AddCommand(upgrade.CmdUpgrade)
	CmdRoot.AddCommand(gen.CmdGen)
	create.CmdCreate.AddCommand(create.CmdCreateHandler)
	create.CmdCreate.AddCommand(create.CmdCreateService)
	create.CmdCreate.AddCommand(create.CmdCreateRepo)
	
	create.CmdCreate.AddCommand(create.CmdCreateAll)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}
