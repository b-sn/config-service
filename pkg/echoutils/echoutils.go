package echoutils

import (
	"configer-service/internal/repository"
	"configer-service/pkg/utils"
	"net"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/sirupsen/logrus"
)

func AuthByIP(IPList []net.IPNet) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			ip := utils.GetIP(ctx.Request())

			if ip.IsLoopback() {
				return next(ctx)
			}

			if len(IPList) > 0 {
				for _, ipnet := range IPList {
					if ipnet.Contains(ip) {
						return next(ctx)
					}
				}
			}

			return ReturnDefault401()
		}
	}
}

func AuthorizeUser(userSecret []byte, userRepo repository.UserI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			verifiedToken, err := jwt.ParseHeader(
				ctx.Request().Header,
				"X-Config-User-Token",
				jwt.WithKey(jwa.HS256, userSecret),
			)
			if err != nil {
				return ReturnDefault401()
			}

			sub := verifiedToken.Subject()
			userUuid, err := uuid.Parse(sub)
			if err != nil {
				logrus.Warnf("authorize user error: unexpected uuid [%s]", sub)
				return ReturnDefault401()
			}

			userID := userUuid.String()
			if activeUser, err := userRepo.ExistsAndActive(userID); err != nil {
				logrus.Errorf("authorize user error: %w", err)
				return ReturnDefault401()

			} else if !activeUser {
				logrus.Infof("inactive user requested: %s", userID)
				return ReturnDefault401()
			}

			ctx.Set("user_id", userID)

			return next(ctx)
		}
	}
}

func AuthorizeSingleUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("user_id", uuid.Nil)
			return next(ctx)
		}
	}
}

func ReturnDefault404() error {
	return echo.NewHTTPError(http.StatusNotFound)
}

func ReturnDefault401() error {
	return echo.NewHTTPError(http.StatusUnauthorized)
}

func ReturnDefault400() error {
	return echo.NewHTTPError(http.StatusBadRequest)
}
