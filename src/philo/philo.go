package philo

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/stRuct"
)

func PhiloMsg(philoStruct *stRuct.PhiloStruct) {
	fmt.Println("This is ", philoStruct.Number, "!")
}
