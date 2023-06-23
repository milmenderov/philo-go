package main

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/internal/app"
	"os"
)

func main() {

	fmt.Println(os.Args)
	app.Run(os.Args)
}
