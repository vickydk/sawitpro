package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) Signup(ctx echo.Context) error {
	request := &generated.SignupRequest{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.ErrorResponse{
			Message: err.Error(),
		})
	}
	userUUID, err := s.createNewUser(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &generated.SignupResponse{
		UserUUID: userUUID,
	})
}

func (s *Server) Login(ctx echo.Context) error {
	request := &generated.LoginRequest{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.ErrorResponse{
			Message: err.Error(),
		})
	}
	resp, err := s.checkAuth(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) MeDetail(ctx echo.Context) error {
	userUUID := fmt.Sprintf("%v", ctx.Get("userUUID"))
	resp, err := s.findUserByID(userUUID)
	if err != nil {
		return ctx.NoContent(http.StatusForbidden)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateUser(ctx echo.Context) error {
	request := &generated.UpdateUserRequest{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.ErrorResponse{
			Message: err.Error(),
		})
	}
	userUUID := fmt.Sprintf("%v", ctx.Get("userUUID"))
	resp, err := s.userUpdate(userUUID, request)
	if err != nil {
		return ctx.NoContent(http.StatusConflict)
	}
	return ctx.JSON(http.StatusOK, resp)
}
