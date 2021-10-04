package checker

import (
	"fmt"
	"net/http"
	"strings"
)

type result struct {
	status string
}

func CheckStatus(url string) result {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		failStatus := fmt.Sprintf("Failed with status %d", res.StatusCode)
		return result{failStatus}
	} else {
		successStatus := fmt.Sprintf("Successful with status %d", res.StatusCode)
		return result{successStatus}
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
