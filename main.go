package main

import (
	"proyeccionesFAMED/app"
)

func main() {

	if err := app.StartBackend(); err != nil {
		panic(err)
	}

}
