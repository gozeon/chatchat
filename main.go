package main

import (
	"chatchat/database"
	"chatchat/model"
	"chatchat/public"
	"encoding/json"
	"fmt"
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
	err := db.Open("chatchat.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	band := model.NewBand(db)
	avatar := model.NewAvatar(db)
	other := model.NewOther(db)
	message := model.NewMessage(db)
	keyword := model.NewKeyword(db)

	// 初始化
	band.Init()
	avatar.Init()
	other.Init()
	message.Init()
	keyword.Init()

	e := echo.New()
	e.Renderer = public.New()

	e.Use(middleware.Timeout())
	e.Use(middleware.Decompress())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.StaticFS("/", public.StaticDir)
	e.Static("/", "public")

	e.Static("/upload", "upload").Name = "Upload"

	e.GET("/", func(c echo.Context) error {
		c.Logger().Print(c.Logger().Level())
		return c.Redirect(302, c.Echo().Reverse("Band"))
	})

	e.GET("/chat", func(c echo.Context) error {
		debug := c.QueryParam("debug")
		return c.Render(http.StatusOK, "chat.html", map[string]interface{}{
			"Band":    band.Get(),
			"Avatar":  avatar.Get(),
			"Other":   other.Get(),
			"Message": message.Get("default"),
			"Quick":   message.Get("quick"),
			"Debug":   debug,
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

	e.GET("/message", func(c echo.Context) error {
		return c.Render(http.StatusOK, "message.html", map[string]interface{}{
			"title":      "Message",
			"ActivePage": "message",
			"Message":    message.Get("default"),
			"Quick":      message.Get("quick"),
		})
	})

	e.GET("/quick", func(c echo.Context) error {
		return c.Render(http.StatusOK, "quick.html", map[string]interface{}{
			"title":      "Quick",
			"ActivePage": "quick",
			"Message":    message.Get("default"),
			"Quick":      message.Get("quick"),
		})
	})

	e.GET("/keyword", func(c echo.Context) error {
		return c.Render(http.StatusOK, "keyword.html", map[string]interface{}{
			"title":      "Keyword",
			"ActivePage": "keyword",
			"Band":       band.Get(),
		})
	})

	e.GET("/keyword-default", func(c echo.Context) error {
		return c.Render(http.StatusOK, "keyword-default.html", map[string]interface{}{
			"title":      "Keyword",
			"ActivePage": "keyword",
			"Message":    keyword.Get("default"),
		})
	})

	e.GET("/keyword-list", func(c echo.Context) error {
		return c.Render(http.StatusOK, "keyword-list.html", map[string]interface{}{
			"title":       "Keyword",
			"ActivePage":  "keyword",
			"KeywordList": keyword.GetList(),
		})
	})

	e.Any("/reply", func(c echo.Context) error {
		var msg string
		var reply map[string]any

		var q struct {
			Q string `json:"q" form:"q" query:"q"`
		}
		err := c.Bind(&q)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		msg = keyword.Get(q.Q).(string)

		fmt.Println()
		if len(msg) == 0 {
			msg = keyword.Get("default").(string)
		}

		json.Unmarshal([]byte(msg), &reply)

		if len(reply) == 0 {
			// 不回复
			return c.JSON(http.StatusOK, nil)
		}

		return c.JSON(http.StatusOK, reply)
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
