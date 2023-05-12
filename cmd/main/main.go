package main

import (
	"fmt"
	"net/http"

	"github.com/b0gochort/conv_service/config"
	appHTTP "github.com/b0gochort/conv_service/internal/delivery/http"
	appMiddleware "github.com/b0gochort/conv_service/internal/delivery/middleware"
	"github.com/b0gochort/conv_service/internal/infrastucture"
	"github.com/b0gochort/conv_service/internal/usecase"
	"github.com/b0gochort/conv_service/internal/utils/crypto"
	"github.com/b0gochort/conv_service/internal/utils/jwt"
	"github.com/b0gochort/conv_service/pkg/postgres"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

// func SignUp(c echo.Context) error {
// 	var req request.SignUpReq
// 	var UserUC usecase.UserUseCase

// 	if err := c.Bind(&req); err != nil {
// 		// h.logger.Error("json not parsed c.Bind/user_handler.SignUp")
// 		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
// 			"message": "json not parsed",
// 		})
// 	}
// 	if err := req.Validate(); err != nil {
// 		errVal := err.(validation.Errors)
// 		// h.logger.Error("Cannot validate json Validate/user_handler.SignUp")
// 		return c.JSON(http.StatusBadRequest, errVal)
// 	}

// 	if err := UserUC.SignUp(&req); err != nil {
// 		// h.logger.Error("SignUp/user_handler.SignUp")
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "signup successfully",
// 	})
// }

func main() {
	// Load config
	config, confErr := config.LoadConfig()
	if confErr != nil {
		panic(confErr)
	}
	print(config.JWTSecretKey.JWTSecretKey)
	// setup logger
	logger, logErr := zap.NewDevelopment()
	if logErr != nil {
		panic(fmt.Sprintf("Error when init logger: %v", logErr))
	}
	// Connetc to DataBase
	db, dbErr := postgres.NewPostgres(postgres.PG{
		Addr:         config.PG.Addr,
		User:         config.PG.User,
		Password:     config.PG.Password,
		DataBaseName: config.PG.DataBaseName,
	})
	if dbErr != nil {
		logger.Panic("Error when connect to DB %s", zap.Field{
			Interface: dbErr,
		})
	}
	// SetUp repository
	userRepo := infrastucture.NewUserRepo(db)

	// Setup Service
	cryptoSvc := crypto.NewCryptoService()
	jwtSvc := jwt.NewJWTService(config.JWTSecretKey.JWTSecretKey)

	// Setup usecase
	userUC := usecase.NewUsersUC(userRepo, cryptoSvc, jwtSvc)

	// Setup app middleware
	appMiddleware := appMiddleware.NewMiddleware(jwtSvc)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(echozap.ZapLogger(logger))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am alive")
	})
	// e.POST("/user/signup", SignUp)
	appHTTP.NewUserHandler(e, appMiddleware, userUC, logger)

	if err := e.Start(config.HTTP.PORT); err != nil {
		panic(err)
	}

	defer db.Close()

}
