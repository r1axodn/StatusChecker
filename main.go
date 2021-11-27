package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckStatus(urlOrIp string) (string, int) {
	res, err := http.Get(urlOrIp)
	if err != nil {
		return "err!", http.StatusInternalServerError
	}
	switch {
	case res.StatusCode >= 400 && res.StatusCode < 500:
		failStatus := fmt.Sprintf("Failed with status %d", res.StatusCode)
		return failStatus, res.StatusCode
	case res.StatusCode >= 200 && res.StatusCode < 300:
		successStatus := fmt.Sprintf("Successful with status %d", res.StatusCode)
		return successStatus, res.StatusCode
	case res.StatusCode >= 300 && res.StatusCode < 400:
		redirectingStatus := "Redirecting maybe"
		return redirectingStatus, res.StatusCode
	case res.StatusCode >= 500:
		StatusInternalServerError := "Oki.. the server has a problem. Give up"
		return StatusInternalServerError, res.StatusCode
	case res.StatusCode >= 100 && res.StatusCode < 200:
		status100 := "Loading...99%"
		return status100, res.StatusCode
	}
	return fmt.Sprintf("Successful with status %d", res.StatusCode), res.StatusCode
}

func handleHome(c echo.Context) error {
	return c.File("./html/home.html")
}

func handleChecker(c echo.Context) error {
	_, status := CheckStatus("144.172.75.157")
	switch {
	case status >= 400 && status < 500:
		return c.File("./html/400.html")
	case status >= 300 && status < 400:
		return c.File("./html/300.html")
	case status >= 100 && status < 200:
		return c.File("./html/100.html")
	case status >= 500:
		return c.File("./html/500.html")

	}
	return c.File("./html/200.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/check", handleChecker)
	e.Logger.Fatal(e.Start(":3000"))
}
