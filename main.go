package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type result struct {
	status string
}

func CheckStatus(url string) (string, int) {
	res, err := http.Get(url)
	if err != nil {
		return "err!", http.StatusInternalServerError
	}
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		failStatus := fmt.Sprintf("Failed with status %d", res.StatusCode)
		return failStatus, res.StatusCode
	} else if res.StatusCode >= 200 && res.StatusCode < 300 {
		successStatus := fmt.Sprintf("Successful with status %d", res.StatusCode)
		return successStatus, res.StatusCode
	} else if res.StatusCode >= 300 && res.StatusCode < 400 {
		status300 := "Hmm... the client is choosing multiple things."
		return status300, res.StatusCode
	} else if res.StatusCode >= 500 {
		StatusInternalServerError := fmt.Sprintf("Oki.. the server has a problem. Give up")
		return StatusInternalServerError, res.StatusCode
	} else if res.StatusCode >= 100 && res.StatusCode < 200 {
		status100 := "Loading...99%"
		return status100, res.StatusCode
	}
	return fmt.Sprintf("Successful with status %d", res.StatusCode), res.StatusCode
}

func handleHome(c echo.Context) error {
	return c.File("./html/home.html")
}

func handleChecker(c echo.Context) error {
	result, status := CheckStatus("https://discord.com/")
	fmt.Println(result)
	if status >= 400 && status < 500 {
		return c.File("./html/400.html")
	} else if status >= 300 && status < 500 {
		return c.File("./html/300.html")
	} else if status >= 500 {
		return c.File("./html/500.html")
	} else if status >= 100 && status < 200 {
		return c.File("./html/100.html")
	}
	return c.File("./html/200.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/check", handleChecker)
	e.Logger.Fatal(e.Start(":3000"))
}
