package utils

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func createLink4Markdown(url string, anchorText string, followRate int) (string, error) {
	if url == "" {
		return "", errors.New("url is empty.")
	}

	at := anchorText
	if anchorText == "" {
		at = "여기"
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	r := rand.Intn(100) + 1

	rel := ""
	if followRate < r {
		rel = "{:rel=\"nofollow\"}"
	}

	return fmt.Sprintf("[%s](%s)%s", at, url, rel), nil
}
