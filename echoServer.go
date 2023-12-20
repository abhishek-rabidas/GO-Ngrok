package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"net/http"
)

func RunServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up!")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up! TEST")
	})

	err := runNgrok(context.Background(), e)

	if err != nil {
		return
	}

}

func runNgrok(ctx context.Context, e *echo.Echo) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)

	if err != nil {
		return err
	}

	log.Println("Tunnel URL", listener.URL())

	return http.Serve(listener, e)

}
