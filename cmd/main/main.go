package main

import (
	"fmt"
	"net/http"

	"github.com/b0gochort/conv_service/pkg/postgres"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Hellohandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

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

	logger, logErr := zap.NewDevelopment()
	if logErr != nil {
		panic(fmt.Sprintf("Error when init logger: %v", logErr))
	}
	db, dbErr := postgres.NewPostgres()
	if dbErr != nil {
		logger.Panic("Error when connect to DB %s", zap.Field{
			Interface: dbErr,
		})
	}
	e := echo.New()
	e.Use(echozap.ZapLogger(logger))
	e.GET("/", Hellohandler)
	if err := e.Start(":8000"); err != nil {
		panic(err)
	}

	defer db.Close()

}
