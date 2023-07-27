package main

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/app"
	"os"
	"strconv"
)

func InvalidArgs() {
	if len(os.Args) < 5 || len(os.Args) > 6 || os.Args[0] == "" {
		fmt.Println("Invalid arguments")
		os.Exit(0)
	}

	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil || num <= 0 {
			fmt.Println("Invalid arguments")
			os.Exit(0)
		}
	}
}

func main() {

	InvalidArgs()

	//fmt.Println(os.Args)
	app.Run(os.Args)

}
