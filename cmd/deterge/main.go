package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	if err := executeCommand(); err != nil {
		log.Fatal(err)
	}
}

func executeCommand() error {
	rootCommand := &cobra.Command{
		Use:   "deterge",
		Short: "Tool for sanitising data files of sensitive information through substitution with fake information.",
		Run: func(cmd *cobra.Command, args []string) {
			// Show help when no command is specified
			_ = cmd.Help()
		},
	}

	rootCommand.AddCommand(Serve{}.Cobra())

	return rootCommand.Execute()
}
