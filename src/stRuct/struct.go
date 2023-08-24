package stRuct

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type TableStruct struct {
	PhiloCount int
	TimeLife   int64
	TimeEat    int64
	TimeSleep  int64
	FullPhilo  int
	MustEat    int
	StartTime  int64
	PrintMutex sync.Mutex
	Philos     []PhiloStruct
	Forks      []sync.Mutex
}

type PhiloStruct struct {
	Number    int
	EatCount  int
	TimeDie   int64
	LeftFork  *sync.Mutex
	RightFork *sync.Mutex
	Table     *TableStruct
}

func PrintAction(philo *PhiloStruct, Action string) {
	timer := Track_time(philo.Table)
	livetime := philo.TimeDie - Track_time(philo.Table)
	philo.Table.PrintMutex.Lock()
	fmt.Println("TIME:", timer, "PhiloLive", livetime, "Philo number:", philo.Number, Action)
	philo.Table.PrintMutex.Unlock()
}

func Track_time(table *TableStruct) int64 {
	time_passed := time.Now().UnixMilli() - table.StartTime
	return time_passed
}

func CheckDead(philo *PhiloStruct) {
	livetime := philo.TimeDie - Track_time(philo.Table)
	if philo.Table.TimeEat >= livetime {
		PrintAction(philo, "eating")
		time.Sleep(time.Duration(philo.Table.TimeEat) * time.Millisecond)
		livetime = philo.TimeDie - Track_time(philo.Table)
		PrintAction(philo, "Dead")
		os.Exit(0)
	} else if philo.Table.TimeSleep >= livetime {
		PrintAction(philo, "sleeping")
		time.Sleep(time.Duration(philo.Table.TimeSleep) * time.Millisecond)
		livetime = philo.TimeDie - Track_time(philo.Table)
		PrintAction(philo, "Dead")
		os.Exit(0)
	}
}

func (philo *PhiloStruct) Eat() {
	if os.Args[1] == "1" {
		philo.LeftFork.Lock()
		PrintAction(philo, "take left fork!")
		philo.LeftFork.Unlock()
		PrintAction(philo, "drop left fork!")
		philo.Sleep()
		PrintAction(philo, "Dead")
		os.Exit(0)
	} else {
		philo.LeftFork.Lock()
		PrintAction(philo, "take left fork!")
		philo.RightFork.Lock()
		PrintAction(philo, "take right fork!")
		CheckDead(philo)
		PrintAction(philo, "eating")
		time.Sleep(time.Duration(philo.Table.TimeEat) * time.Millisecond)

		philo.EatCount++
		philo.TimeDie = Track_time(philo.Table) + philo.Table.TimeLife

		philo.RightFork.Unlock()
		PrintAction(philo, "drop right fork!")
		philo.LeftFork.Unlock()
		PrintAction(philo, "drop left fork!")
	}
}

func (philo *PhiloStruct) Sleep() {
	CheckDead(philo)
	PrintAction(philo, "sleeping")
	time.Sleep(time.Duration(philo.Table.TimeSleep) * time.Millisecond)
	PrintAction(philo, "thinking")
}
