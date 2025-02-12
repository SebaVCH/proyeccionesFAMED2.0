package app

import (
	"proyeccionesFAMED/config"
	"proyeccionesFAMED/database"
)

func StartBackend() error {

	if err := config.LoadENV(); err != nil {
		return err
	}

	if err := database.StartDB(); err != nil {
		return err
	}

	return nil
}
