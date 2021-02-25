package main

import (
	"github.com/spf13/cobra"
	"github.com/xushuhui/goal/internal/project"
	"log"
)

var (
	// Version is the version of the compiled software.
	Version string = "v0.0.1.dev"

	rootCmd = &cobra.Command{
		Use:     "goal",
		Short:   "Goal: An elegant toolkit for Go Web Developer.",
		Long:    `Goal: An elegant toolkit for Go Web Developer.`,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(project.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
