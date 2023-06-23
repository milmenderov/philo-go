package app

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/internal/models"
)

func Run(args []string) {
	fmt.Println("Run initialization whole app")
	var mainTable = new(models.TableStruct)
	//mainTable.PhiloCount = 1
	//println(mainTable.PhiloCount)
	parseArgs(args, mainTable)
	//println(mainTable.PhiloCount)
}

func parseArgs(args []string, mainTable *models.TableStruct) {
	fmt.Println("Parsing args")
	//mainTable.PhiloCount = 10
}
