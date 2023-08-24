package app

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/stRuct"
	"gitlab.com/e.ilmenderov/philo-go/src/table"
	"time"
)

func even_philo(philo_tables *stRuct.TableStruct) {
	for i := 0; i < philo_tables.PhiloCount; i++ {
		philo_tables.Philos[i].Number = i
		philo_tables.Philos[i].Table = philo_tables
		if i%2 == 0 {
			philo_tables.Philos[i].LeftFork = &philo_tables.Forks[i]
			if i+1 == philo_tables.PhiloCount {
				philo_tables.Philos[i].RightFork = &philo_tables.Forks[0]
			} else {
				philo_tables.Philos[i].RightFork = &philo_tables.Forks[i+1]
			}
		}
	}
}

func odd_philo(philo_tables *stRuct.TableStruct) {
	for i := 0; i < philo_tables.PhiloCount; i++ {
		if i%2 != 0 {
			philo_tables.Philos[i].LeftFork = &philo_tables.Forks[i-1]
			philo_tables.Philos[i].RightFork = &philo_tables.Forks[i]
		}
	}
}

func life(philo *stRuct.PhiloStruct) {
	philo.TimeDie = stRuct.Track_time(philo.Table) + philo.Table.TimeLife
	for {
		if philo.Table.MustEat == philo.EatCount {
			philo.Table.FullPhilo++
			break
		}
		philo.Eat()
		philo.Sleep()
	}
}

func Run(args []string) {
	fmt.Println("Run initialization whole app")
	philo_tables := table.ParseArgs(args)
	even_philo(philo_tables)
	odd_philo(philo_tables)
	philo_tables.StartTime = time.Now().UnixMilli()
	for i := 0; i < philo_tables.PhiloCount; i++ {
		go life(&philo_tables.Philos[i])
	}
	for {
		if philo_tables.FullPhilo == philo_tables.PhiloCount {
			break
		}
	}
}
