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
			//fmt.Println("Philo: ", i)
			philo_tables.Philos[i].LeftFork = &philo_tables.Forks[i]
			//fmt.Println("lf - ", i)
			if i+1 == philo_tables.PhiloCount {
				philo_tables.Philos[i].RightFork = &philo_tables.Forks[0]
				//fmt.Println("rf - 0")
			} else {
				philo_tables.Philos[i].RightFork = &philo_tables.Forks[i+1]
			}
		}
		//0 1 2 3 4 5 total - 6
		//0 - 5р и 0л
		//1 - 1л и 0р
		//2 - 1р и 2л
		//3 - 3л и 2р
		//4 - 3р и 4л
		//5 - 5л и 4р
	}
}
func odd_philo(philo_tables *stRuct.TableStruct) {
	for i := 0; i < philo_tables.PhiloCount; i++ {
		if i%2 != 0 {
			//fmt.Println("Philo: ", i)
			philo_tables.Philos[i].LeftFork = &philo_tables.Forks[i-1]
			//fmt.Println("lf - ", i-1)
			philo_tables.Philos[i].RightFork = &philo_tables.Forks[i]
			//fmt.Println("rf - ", i)
			//0 1 2 3 4 5 total - 6

			//0 - 5р и 0л
			//1 - 1л и 0р
			//2 - 1р и 2л
			//3 - 3л и 2р
			//4 - 3р и 4л
			//5 - 5л и 4р
		}
	}
}

func life(philo *stRuct.PhiloStruct) {
	for {
		if philo.Table.MustEat == philo.EatCount {
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
		if philo_tables.MustEat == philo_tables.Philos[0].EatCount {
			break
		}
	}

}
