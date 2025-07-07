package utils

import (
	"fmt"
	"net/http"
	"time"
)

func HTTPWithRetry(f func(string) (*http.Response, error), url string) (*http.Response, error) {
	var resp *http.Response
	var err error
	count := 10

	for i := 0; i < count; i++ {
		resp, err = f(url)
		if err != nil {
			fmt.Printf("Error calling url %v\n", url)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	return resp, err
}
