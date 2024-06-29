package echo

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/routes"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoServer struct {
	router      *echo.Echo
	repoHandler database.RepositoryHandler
}

func NewEchoServer(repoHandler database.RepositoryHandler) *echoServer {
	return &echoServer{
		router:      echo.New(),
		repoHandler: repoHandler,
	}
}

func (s *echoServer) Listen() {
	s.setAppHandlers(s.router)

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())

	// s.router.Server = &http.Server{
	// 	Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 15 * time.Second,
	// }

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.router.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.router.Shutdown(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Server stopped")
}

func (s *echoServer) setAppHandlers(router *echo.Echo) {
	api := router.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	routes.LoadUserRoutes(api, s.repoHandler)

}
