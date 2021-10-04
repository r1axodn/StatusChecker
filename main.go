package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type result struct {
	status string
}

func CheckStatus(url string) string {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		/*
			return result{failStatus}*/
		failStatus := fmt.Sprintf("Failed with status %d", res.StatusCode)
		return failStatus
	} else {
		/*
			successStatus := fmt.Sprintf("Successful with status %d", res.StatusCode)
			return result{successStatus}*/
		successStatus := fmt.Sprintf("Successful with status %d", res.StatusCode)
		return successStatus
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleChecker(c echo.Context) error {
	url := strings.ToLower(CleanString(c.FormValue("urlInput")))
	status := CheckStatus(url)
	fmt.Println(status)
	return c.String(http.StatusOK, status)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/check", handleChecker)
	e.Logger.Fatal(e.Start(":3000"))
}
