package http

import (
	"net/http"

	"github.com/b0gochort/conv_service/internal/delivery/middleware"
	"github.com/b0gochort/conv_service/internal/transport/request"
	"github.com/b0gochort/conv_service/internal/usecase"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserUC usecase.UserUseCase
	logger zap.Logger
}

func NewUserHandler(e *echo.Echo, middleware *middleware.Middleware, userUC usecase.UserUseCase, logger *zap.Logger) {
	handler := &UserHandler{
		UserUC: userUC,
		logger: *logger,
	}
	handler.logger.Info("Registrate user routes")
	apiV1 := e.Group("/api/v1")
	apiV1.POST("/user/signup", handler.SignUp)
	apiV1.POST("/user/login", handler.LogIn)

}

func (h *UserHandler) SignUp(c echo.Context) error {
	var req request.SignUpReq

	if err := c.Bind(&req); err != nil {
		h.logger.Error("json not parsed c.Bind/user_handler.SignUp")
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "json not parsed",
		})
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		h.logger.Error("Cannot validate json Validate/user_handler.SignUp")
		return c.JSON(http.StatusBadRequest, errVal)
	}

	if err := h.UserUC.SignUp(&req); err != nil {
		h.logger.Error("SignUp/user_handler.SignUp")
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "signup successfully",
	})
}
func (h *UserHandler) LogIn(c echo.Context) error {

	var req request.LogInReq

	if err := c.Bind(&req); err != nil {
		h.logger.Error("json not parsed c.Bind/user_handler.LogIn")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := req.Validate(); err != nil {
		h.logger.Error("Cannot validate json Validate/user_handler.LogIn")
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, errVal)
	}

	accessToken, err := h.UserUC.LogIn(&req)

	if err != nil {
		h.logger.Error("user not found LigIn/user_handler.LogIn")
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"access_token": accessToken,
		},
	})
}
