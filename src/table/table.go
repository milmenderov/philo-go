package table

import (
	"gitlab.com/e.ilmenderov/philo-go/src/stRuct"
	"math"
	"strconv"
	"sync"
)

func ParseArgs(args []string) *stRuct.TableStruct {

	table := new(stRuct.TableStruct)
	table.PhiloCount, _ = strconv.Atoi(args[1])
	tmp, _ := strconv.Atoi(args[2])
	table.TimeLife = int64(tmp)
	tmp, _ = strconv.Atoi(args[3])
	table.TimeEat = int64(tmp)
	tmp, _ = strconv.Atoi(args[4])
	table.TimeSleep = int64(tmp)
	if len(args) == 6 {
		table.MustEat, _ = strconv.Atoi(args[5])
	} else {
		table.MustEat = math.MaxInt64
	}
	table.Forks = append(table.Forks, make([]sync.Mutex, table.PhiloCount)...)
	table.Philos = append(table.Philos, make([]stRuct.PhiloStruct, table.PhiloCount)...)
	for i := 0; i < table.PhiloCount; i++ {
		table.Philos[i].Number = i
	}
	return table
}
