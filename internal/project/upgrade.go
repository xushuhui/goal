package project

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xushuhui/goal/internal/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the goal tools",
	Long:  "Upgrade the goal tools. Example: goal upgrade",
	Run:   Run,
}

// Run upgrade the goal tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoGet(
		"github.com/go-kratos/kratos/cmd/kratos/v2",
	)
	if err != nil {
		fmt.Println(err)
	}
}