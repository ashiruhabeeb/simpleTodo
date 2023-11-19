package router

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashiruhabeeb/simpleTodoApp/handler"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/repository"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, port string, db *sql.DB) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	logger := logger.NewSlogHandler()

	todoRepo := repository.NewTodoRepo(db)
	todoHandler :=  handler.NewTodoService(todoRepo, logger)

	e.POST("/signup", todoHandler.Store)

	httpSrv := &http.Server{
		Addr:           ":" + port,
		Handler:        e,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		IdleTimeout:    time.Second * 10,
	}

	go func() {
		logger.Info("[INIT] ✅ gin router running and listening on port")

		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("[ERROR] http.ListenAndServe failure: %v\n", err)
		}
	}()

	// declare a buffered channel that reveives unix signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	logger.Warn("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		logger.Warn("[CLOSE] error when shutdown server: %s", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Warn("[CLOSE] timeout of 5 seconds.")
	logger.Warn("[CLOSE] server exiting")
}
