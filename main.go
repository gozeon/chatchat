package main

import (
	"chatchat/database"
	"chatchat/model"
	"chatchat/public"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = os.MkdirAll("upload", 0755)
	if err != nil {
		log.Fatal("Error create upload folder")
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
	avatar := model.NewAvatar(db)
	other := model.NewOther(db)

	// 初始化
	band.Init()
	avatar.Init()
	other.Init()

	e := echo.New()
	e.Renderer = public.New()

	e.Use(middleware.Timeout())
	e.Use(middleware.Decompress())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.StaticFS("/", public.StaticDir)
	e.Static("/upload", "upload").Name = "Upload"

	e.GET("/", func(c echo.Context) error {
		c.Logger().Print(c.Logger().Level())
		return c.Redirect(302, c.Echo().Reverse("Band"))
	})

	e.GET("/chat", func(c echo.Context) error {
		debug := c.QueryParam("debug")
		return c.Render(http.StatusOK, "chat.html", map[string]interface{}{
			"Band":   band.Get(),
			"Avatar": avatar.Get(),
			"Other":  other.Get(),
			"Debug":  debug,
		})
	})

	e.GET("/band", func(c echo.Context) error {
		return c.Render(http.StatusOK, "band.html", map[string]interface{}{
			"title":      "Band",
			"ActivePage": "band",
			"Band":       band.Get(),
		})
	}).Name = "Band"

	e.GET("/avatar", func(c echo.Context) error {
		return c.Render(http.StatusOK, "avatar.html", map[string]interface{}{
			"title":      "Avatar",
			"ActivePage": "avatar",
			"Avatar":     avatar.Get(),
		})
	}).Name = "Avatar"

	e.GET("/other", func(c echo.Context) error {
		return c.Render(http.StatusOK, "other.html", map[string]interface{}{
			"title":      "Other",
			"ActivePage": "other",
			"Other":      other.Get(),
		})
	}).Name = "Other"

	e.GET("/link", func(c echo.Context) error {
		return c.Render(http.StatusOK, "link.html", nil)
	})

	e.POST("/dbsave", func(c echo.Context) (err error) {
		bucketName := c.FormValue("bucketName")
		key := c.FormValue("key")
		value := c.FormValue("value")

		if err = db.Put(bucketName, []byte(key), []byte(value)); err != nil {
			c.Logger().Error(err)
			return
		}

		return c.NoContent(http.StatusOK)
	})

	e.POST("/upload", func(c echo.Context) error {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		id, _ := uuid.NewUUID()
		p := filepath.Join("upload", id.String()+filepath.Ext(file.Filename))
		// Destination
		dst, err := os.Create(p)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.String(http.StatusOK, filepath.ToSlash(p))
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
