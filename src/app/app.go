package app

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/table"
)

func Run(args []string) {
	fmt.Println("Run initialization whole app")
	table.ParseArgs(args)
}
