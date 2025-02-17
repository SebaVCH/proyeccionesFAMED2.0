package app

import (
	"proyeccionesFAMED/config"
	"proyeccionesFAMED/database"
	"proyeccionesFAMED/routes"
)

func StartBackend() error {

	if err := config.LoadENV(); err != nil {
		return err
	}

	if err := database.StartDB(); err != nil {
		return err
	}

	if err := routes.SetupRouter().Run(":8080"); err != nil {
		return err
	}

	return nil
}
