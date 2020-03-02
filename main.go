package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// ルーティング
	e.GET("/hello", mainPage)

	// サーバー起動
	e.Start(":8080")
}

func mainPage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

// func mainPage() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
// 	}
// }
