package main

import (
	"chatchat/public"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	e := echo.New()
	e.Renderer = public.New()

	e.Use(middleware.Timeout())
	e.Use(middleware.Decompress())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.StaticFS("/", public.StaticDir)

	e.GET("/", func(c echo.Context) error {
		c.Logger().Print(c.Logger().Level())
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"title": "Home",
		})
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
