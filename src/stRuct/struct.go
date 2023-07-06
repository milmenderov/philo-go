package stRuct

import (
	"fmt"
	"time"
)

type TableStruct struct {
	PhiloCount int
	TimeDie    int
	TimeEat    int
	TimeSleep  int
	MustEat    int
	FullPhilo  int
	StartTime  int64
	Philos     []PhiloStruct
	Forks      []bool
}

type PhiloStruct struct {
	Number    int
	State     int
	EatCount  int
	LeftFork  *bool
	RightFork *bool
	Table     *TableStruct
}

func track_time(table *TableStruct) int64 {
	time_passed := time.Now().UnixMilli() - table.StartTime
	return time_passed
}

func (philo *PhiloStruct) Eat() {

	for *philo.LeftFork != false {
	}
	*philo.LeftFork = true
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "take left fork!", philo.LeftFork)
	for *philo.RightFork != false {
	}
	*philo.RightFork = true
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "take right fork!", philo.RightFork)
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "eat")
	time.Sleep(time.Duration(philo.Table.TimeEat) * time.Millisecond)
	philo.EatCount++
	*philo.RightFork = false
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "drop right fork!", philo.RightFork)
	*philo.LeftFork = false
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "drop left fork!", philo.LeftFork)
}

func (philo *PhiloStruct) Sleep() {
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "sleep")
	time.Sleep(time.Duration(philo.Table.TimeSleep) * time.Millisecond)
	fmt.Println("TIME:", track_time(philo.Table), "Philo number:", philo.Number, "thinking")
}
