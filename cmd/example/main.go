// Package main handles execution of the example cmd process
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"

	"github.com/bkonkle/go-garbage/internal/example"
	"github.com/bkonkle/go-garbage/internal/example/handlers"
	"github.com/bkonkle/go-garbage/internal/example/io"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	app := fizz.NewFromEngine(engine)

	infos := &openapi.Info{
		Title:       "Go Garbage",
		Description: "A Go garbage collection example",
		Version:     "v0.1.0",
	}

	app.GET("/openapi.json", nil, app.OpenAPI(infos, "json"))
	app.GET("/docs", nil, func(c *gin.Context) {
		c.HTML(http.StatusOK, "swagger.html", gin.H{
			"openapi_url": "/openapi.json",
			"title":       "Go Garbage",
		})
	})

	group := app.Group("", "garbage", "Garbage Collection")

	group.GET("/allocate-memory", []fizz.OperationOption{
		fizz.Summary("Allocate Memory"),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", io.HTTPError{}, nil, nil),
	}, tonic.Handler(handlers.Allocate, http.StatusOK))

	group.GET("/run-gc", []fizz.OperationOption{
		fizz.Summary("Run Garbage Collector"),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", io.HTTPError{}, nil, nil),
	}, tonic.Handler(handlers.RunGC, http.StatusOK))

	group.GET("/allocate-memory-and-run-gc", []fizz.OperationOption{
		fizz.Summary("Allocate Memory & Run Garbage Collector"),
		fizz.Response(fmt.Sprint(http.StatusInternalServerError), "Server Error", io.HTTPError{}, nil, nil),
	}, tonic.Handler(handlers.AllocateAndRunGC, http.StatusOK))

	if len(app.Errors()) != 0 {
		log.Fatalf("Fizz errors: %v", app.Errors())
	}

	tonic.SetErrorHook(example.ErrHook)

	srv := &http.Server{
		Addr:              ":8000",
		Handler:           http.TimeoutHandler(app, 30*time.Second, "The request timed out\n"),
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Garbage API is now running at: http://localhost:8000")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
