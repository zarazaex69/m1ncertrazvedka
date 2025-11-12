package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	jsFileUrl := "https://gu-st.ru/landings-st/main.4978f1b2316e002d.js"

	resp, err := http.Get(jsFileUrl)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	content := string(body)

	re := regexp.MustCompile(`['"]([^'"]+\.zip)['"]`)
	matches := re.FindAllStringSubmatch(content, -1)

	if len(matches) == 0 {
		fmt.Println("No .zip URLs found inside the script.")
		return
	}

	fmt.Println("Found URLs:")
	for _, match := range matches {
		if len(match) > 1 {
			fmt.Println(match[1])
		}
	}
}
