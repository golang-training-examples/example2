// Package server provides a simple HTTP server
package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-training-examples/example2/internal/pets"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// StatusResponse is a response for /status endpoint
type StatusResponse struct {
	// Hostname is a hostname of the server
	Hostname string `json:"hostname"`

	// StartedTimestamp is a timestamp when server was started
	StartedTimestamp int64 `json:"started_timestamp"`

	// Uptime is a number of seconds since server was started
	Uptime int64 `json:"uptime"`
}

// Server starts HTTP server
func Server(port int, dsn string) {
	db, _ := pets.NewDB(dsn)
	db.Migrate()

	l := log.New("echo")
	l.DisableColor()
	l.SetLevel(log.INFO)

	e := echo.New()
	e.Logger = l
	e.HideBanner = true
	e.HidePort = true

	e.Use(echoprometheus.NewMiddleware("example2"))
	e.GET("/metrics", echoprometheus.NewHandler())

	HOSTNAME, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	STARTED := time.Now().Unix()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<center><h1>Hello World!</h1></center>\n")
	})

	e.GET("/pets", func(c echo.Context) error {
		pets := db.ListPets()
		return c.JSON(http.StatusOK, pets)
	})

	e.POST("/pets", func(c echo.Context) error {
		pet := pets.Pet{}
		c.Bind(&pet)
		db.CreatePet(pet.Name, pet.Age, pet.Kind)
		return c.JSON(http.StatusCreated, pet)
	})

	e.GET("/livez", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK\n")
	})

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, StatusResponse{
			Hostname:         HOSTNAME,
			StartedTimestamp: STARTED,
			Uptime:           time.Now().Unix() - STARTED,
		})
	})

	e.Logger.Infof("Listen on 0.0.0.0:%d, see http://127.0.0.1:%d", port, port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
