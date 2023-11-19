package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, port string) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	httpSrv := &http.Server{
		Addr:           ":" + port,
		Handler:        e,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		IdleTimeout:    time.Second * 10,
	}

	go func() {
		log.Printf("[INIT] ✅ gin router running and listening on port %v", port)

		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] http.ListenAndServe failure: %v\n", err)
		}
	}()

	// declare a buffered channel that reveives unix signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	log.Println("[CLOSE] shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Printf("[CLOSE] error when shutdown server: %s", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("[CLOSE] timeout of 5 seconds.")
	log.Println("[CLOSE] server exiting")
}
