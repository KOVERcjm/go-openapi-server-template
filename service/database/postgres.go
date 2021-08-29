package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kovercheng/driver"
	. "kovercheng/middleware"
	"kovercheng/model/table"
)

var postgres = driver.GetConnection().Postgres

func pgErrorHandler(err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Logger.Infof("No data match in Postgres postgres: %s", err)
	} else {
		Logger.Warnf("Postgres query ERROR: %s", err)
	}
}

func TestPostgres() error {
	if result := postgres.Create(&table.Test{Key: "key", Value: 123}); result.Error != nil {
		pgErrorHandler(result.Error)
		return fmt.Errorf("postgres query error")
	}
	Logger.Infoln("Postgres insert success.")

	if result := postgres.Model(&table.Test{}).Where("key = ?", "key").Update("Value", 456); result.Error != nil {
		pgErrorHandler(result.Error)
		return fmt.Errorf("postgres query error")
	}
	Logger.Infoln("Postgres update success.")

	var findResult table.Test
	if result := postgres.Where("key = ?", "key").First(&findResult); result.Error != nil {
		pgErrorHandler(result.Error)
		return fmt.Errorf("postgres query error")
	}
	Logger.Infof("Postgres find success: %+v", findResult)

	if result := postgres.Where("key = ?", "key").Delete(&table.Test{}); result.Error != nil {
		pgErrorHandler(result.Error)
		return fmt.Errorf("postgres query error")
	}
	Logger.Infoln("Postgres delete success.\n")

	return nil
}
