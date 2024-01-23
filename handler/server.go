package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
)

type Server struct {
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
	}
}

func CreateMiddleware() ([]echo.MiddlewareFunc, error) {
	spec, err := generated.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	validator := middleware.OapiRequestValidatorWithOptions(spec,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
		})

	return []echo.MiddlewareFunc{validator, AuthenticationMiddleware}, nil
}

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.String() == "/login" ||
			c.Request().URL.String() == "/signup" {
			return next(c)
		}
		// Call AuthenticationFunc
		authHeader := c.Request().Header.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return echo.ErrForbidden
		} else {
			tokenString := authHeaderParts[1]
			claimsToken := &claims{}
			token, _ := jwt.ParseWithClaims(tokenString, claimsToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(signingSecret), nil
			})

			if token == nil {
				return echo.ErrUnauthorized
			}

			if token.Valid && !checkTokenExpiry(claimsToken.StandardClaims.ExpiresAt) {
				c.Set("userUUID", claimsToken.UserUUID)
			} else {
				return echo.ErrUnauthorized
			}
		}

		// Continue to the next middleware or route handler
		return next(c)
	}
}

func checkTokenExpiry(timestamp interface{}) bool {
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(time.Now())
		if remainder > 0 {
			return false
		}
	}
	return true
}
