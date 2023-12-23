package cmd

import (
	"fmt"
	"os"
	"todo_list/internal/runner"
)

func Run(args []string) {

	if len(args) < 1 {
		fmt.Println("At least one argument must be entered.")
		os.Exit(1)
	}

	// TODO: Check if runner_type exists.

	runnerType := args[1]
	factory := runner.Factory{Name: runnerType}
	factory.Serve(runnerType).Run(args)
}
