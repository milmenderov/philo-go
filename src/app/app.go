package app

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/table"
)

func Run(args []string) {
	fmt.Println("Run initialization whole app")
	var mainTable = new(table.TableStruct)
	//mainTable.PhiloCount = 1
	//println(mainTable.PhiloCount)
	parseArgs(args, mainTable)
	//println(mainTable.PhiloCount)
}

func parseArgs(args []string, mainTable *table.TableStruct) {
	fmt.Println("Parsing args")
	//mainTable.PhiloCount = 10
}
