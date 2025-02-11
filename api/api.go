package api

import (
	"context"
	"net/http"
	"os"
	"syscall"

	_ "github.com/contact/api/docs"
	"github.com/contact/api/handlers/contact"
	"github.com/contact/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// NewAPI creates a new API
func NewAPI(svc Services) *API {
	return &API{
		server:   echo.New(),
		services: svc,
	}
}

// Start starts the API
func (a *API) Start(ctx context.Context, quit chan os.Signal) {
	cfg := config.Instance()
	a.setupAuthRouter()

	a.setupMiddleware()

	a.setupRoutes()

	defer func() {
		<-ctx.Done()
		quit <- syscall.SIGTERM
	}()

	err := a.server.Start(cfg.Server.Port)
	if err != nil {
		quit <- syscall.SIGTERM
	}
}

// setupMiddleware sets up the middleware
func (a *API) setupMiddleware() {
	cfg := config.Instance()

	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())
	a.server.Use(middleware.CORS())
	a.server.Use(middleware.Secure())
	a.server.Use(middleware.RequestID())

	a.server.Server.ReadTimeout = cfg.Server.Timeout
	a.server.Logger.SetLevel(lvl(cfg.Server.LogLevel))

	// group := a.server.Group(cfg.Server.BaseAddr)
	// config := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(authsvc.JwtCustomClaims)
	// 	},
	// 	SigningKey: authsvc.Secret,
	// }

	// group.Use(echojwt.WithConfig(config))

	// return group
}

// Stop stops the API
func (a *API) Stop(ctx context.Context) error {
	if a.server != nil {
		log.Info("shutting down the server...")
		if err := a.server.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *API) setupAuthRouter() {
	a.server.GET("/swagger/*", echoSwagger.WrapHandler)
	a.server.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
}

func (a *API) setupRoutes() {
	contact.SetupHandler(a.server, a.services.ContactSvc)
}
