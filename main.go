package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	port = flag.String("port", "8080", "port")
)

func main() {
	flag.Parse()

	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "server running")
	})

	router.GET("/secret/:secret", func(c echo.Context) error {
		dbUser, err := ioutil.ReadFile("/run/secrets/db_user")
		if err != nil {
			log.Fatalf("Failed to read db_user secret: %s", err)
		}

		dbPassword, err := ioutil.ReadFile("/run/secrets/db_password")
		if err != nil {
			log.Fatalf("Failed to read db_password secret: %s", err)
		}

		dbhost, err := ioutil.ReadFile("/run/secrets/db_host")
		if err != nil {
			log.Fatalf("Failed to read dbhost secret: %s", err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user":     strings.TrimSuffix(string(dbUser), "\n"),
			"password": strings.TrimSuffix(string(dbPassword), "\n"),
			"host":     strings.TrimSuffix(string(dbhost), "\n"),
		})
	})

	router.Start(":" + *port)
}
