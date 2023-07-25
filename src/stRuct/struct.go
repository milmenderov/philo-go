package stRuct

import (
	"fmt"
	"sync"
	"time"
)

type TableStruct struct {
	PhiloCount int
	TimeDie    int64
	TimeEat    int64
	TimeSleep  int64
	MustEat    int
	FullPhilo  int
	StartTime  int64
	Philos     []PhiloStruct
	Forks      []sync.Mutex
}

type PhiloStruct struct {
	Number    int
	State     int
	EatCount  int
	LeftFork  *sync.Mutex
	RightFork *sync.Mutex
	Table     *TableStruct
}

func track_time(table *TableStruct) int64 {
	time_passed := time.Now().UnixMilli() - table.StartTime
	return time_passed
}

func (philo *PhiloStruct) Eat() {

	philo.LeftFork.Lock()
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "take left fork!")

	philo.RightFork.Lock()
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "take right fork!")

	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "eat")
	time.Sleep(time.Duration(philo.Table.TimeEat) * time.Millisecond)
	philo.EatCount++

	philo.RightFork.Unlock()
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "drop right fork!")

	philo.LeftFork.Unlock()
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "drop left fork!")
}

func (philo *PhiloStruct) Sleep() {
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "sleep")
	time.Sleep(time.Duration(philo.Table.TimeSleep) * time.Millisecond)
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "thinking")
}
