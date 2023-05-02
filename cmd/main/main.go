package main

import (
	"fmt"

	"github.com/b0gochort/conv_service/config"
	"github.com/b0gochort/conv_service/pkg/postgres"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// func versionHandler(db *pg.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var version string
// 		_, err := db.QueryOne(pg.Scan(&version), "SELECT version()")
// 		if err != nil {
// 			return err
// 		}
// 		return c.String(http.StatusOK, version)
// 	}
// }

func main() {

	config, confErr := config.LoadConfig()
	if confErr != nil {
		panic(confErr)
	}

	logger, logErr := zap.NewDevelopment()
	if logErr != nil {
		panic(fmt.Sprintf("Error when init logger: %v", logErr))
	}
	db, dbErr := postgres.NewPostgres(postgres.PG{
		Addr:         config.DataBase.Addr,
		User:         config.DataBase.User,
		Password:     config.DataBase.Password,
		DataBaseName: config.DataBase.DataBaseName,
	})
	if dbErr != nil {
		logger.Panic("Error when connect to DB %s", zap.Field{
			Interface: dbErr,
		})
	}
	e := echo.New()
	e.Use(echozap.ZapLogger(logger))
	if err := e.Start(":8000"); err != nil {
		panic(err)
	}

	defer db.Close()

}
