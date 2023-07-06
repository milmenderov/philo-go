package table

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/stRuct"
	"strconv"
)

func TableMsg(tableStruct *stRuct.TableStruct) {
	fmt.Println("This table have ", tableStruct.PhiloCount, " philosophers")
}

func ParseArgs(args []string) *stRuct.TableStruct {

	table := new(stRuct.TableStruct)
	table.PhiloCount, _ = strconv.Atoi(args[1])
	table.TimeDie, _ = strconv.Atoi(args[2])
	table.TimeEat, _ = strconv.Atoi(args[3])
	table.TimeSleep, _ = strconv.Atoi(args[4])
	if len(args) == 6 {
		table.MustEat, _ = strconv.Atoi(args[5])
	}
	table.Forks = append(table.Forks, make([]bool, table.PhiloCount)...)
	table.Philos = append(table.Philos, make([]stRuct.PhiloStruct, table.PhiloCount)...)
	for i := 0; i < table.PhiloCount; i++ {
		table.Philos[i].Number = i
	}
	return table
}
