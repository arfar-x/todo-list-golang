package runner

import "fmt"

type Cli struct {
	Factory
}

func (c *Cli) GetName() string {
	return "Command-line interface"
}

func (c *Cli) Run([]string) {
	fmt.Println("CLI action")
}
