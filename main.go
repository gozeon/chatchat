package main

import (
	"chatchat/database"
	"chatchat/model"
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
	// 创建并打开数据库
	db := &database.BoltDB{}
	err := db.Open("example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	band := model.NewBand(db)

	// 初始化
	band.Init()

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
		return c.Redirect(302, c.Echo().Reverse("Band"))
	})
	e.GET("/chat", func(c echo.Context) error {
		return c.Render(http.StatusOK, "chat.html", map[string]interface{}{
			"Band": band.Get(),
		})
	})
	e.GET("/band", func(c echo.Context) error {

		return c.Render(http.StatusOK, "band.html", map[string]interface{}{
			"title":      "Band",
			"ActivePage": "band",
		})
	}).Name = "Band"

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
