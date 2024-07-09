package echo

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casbin/casbin"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/middlewares"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/routes"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoServer struct {
	router       *echo.Echo
	logger       logger.Logger
	repoHandler  database.RepositoryHandler
	authEnforcer *casbin.Enforcer
}

func NewEchoServer(repoHandler database.RepositoryHandler, enforcer *casbin.Enforcer, logger logger.Logger) *echoServer {
	return &echoServer{
		router:       echo.New(),
		repoHandler:  repoHandler,
		logger:       logger,
		authEnforcer: enforcer,
	}
}

func (s *echoServer) Listen() {
	s.setAppHandlers(s.router)

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middlewares.UserAuthorizer(s.authEnforcer, s.repoHandler.UserRepository()))

	// s.router.Server = &http.Server{
	// 	Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 15 * time.Second,
	// }

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.router.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
			s.logger.WithError(err).Fatalln("error starting server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.router.Shutdown(ctx); err != nil {
		s.logger.WithError(err).Fatalln("error stopping server")
	}

	s.logger.InfoF("server stopped")
}

func (s *echoServer) setAppHandlers(router *echo.Echo) {
	api := router.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	routes.LoadUserRoutes(api, s.repoHandler)
	routes.LoadAuthRoutes(api, s.repoHandler)
}
