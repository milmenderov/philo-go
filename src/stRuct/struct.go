package stRuct

import (
	"fmt"
	"sync"
	"time"
)

type TableStruct struct {
	PhiloCount int
	TimeLife   int64
	TimeEat    int64
	TimeSleep  int64
	MustEat    int
	FullPhilo  int
	StartTime  int64
	PrintMutex sync.Mutex
	Philos     []PhiloStruct
	Forks      []sync.Mutex
}

type PhiloStruct struct {
	Number    int
	State     int
	EatCount  int
	TimeDie   int64
	LeftFork  *sync.Mutex
	RightFork *sync.Mutex
	Table     *TableStruct
}

func PrintAction(philo *PhiloStruct, Action string) {
	timer := track_time(philo.Table)
	philo.Table.PrintMutex.Lock()
	fmt.Println("TIME:", timer, "Philo number:", philo.Number, Action)
	philo.Table.PrintMutex.Unlock()
}

func track_time(table *TableStruct) int64 {
	time_passed := time.Now().UnixMilli() - table.StartTime
	return time_passed
}

func CheckDeadEat(philo *PhiloStruct) bool {
	if philo.TimeDie > track_time(philo.Table) {
		return false
	}
	return true
}

//func CheckDeadSleep(philo *PhiloStruct) bool {
//	if philo.TimeDie > track_time(philo.Table) {
//		return false
//	}
//	return true
//}

func (philo *PhiloStruct) Eat() {

	philo.LeftFork.Lock()
	PrintAction(philo, "take left fork!")
	philo.RightFork.Lock()
	PrintAction(philo, "take right fork!")
	// чек на смерть
	PrintAction(philo, "eating")
	time.Sleep(time.Duration(philo.Table.TimeEat) * time.Millisecond)
	philo.TimeDie = track_time(philo.Table) + philo.Table.TimeLife
	philo.EatCount++

	philo.RightFork.Unlock()
	PrintAction(philo, "drop right fork!")
	philo.LeftFork.Unlock()
	PrintAction(philo, "drop left fork!")
}

func (philo *PhiloStruct) Sleep() {
	// чек на смерть
	PrintAction(philo, "sleeping")
	time.Sleep(time.Duration(philo.Table.TimeSleep) * time.Millisecond)
	PrintAction(philo, "thinking")
}
