package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type WebsiteResponse struct {
	StatusCode int
	Body       string
	ReceivedAt time.Time
}

func main() {
	// I want to make an application that will repeatedly check the status of a website and keep track of the status codes over time.

	url := "https://www.google.com"
	var responses []WebsiteResponse

	for i:=0; i<10; i++ {
		resp, err := getWebsiteResponse(url)
		if err != nil {
			log.Println("Failed to get response: ", err)
		}
		responses = append(responses, resp)
		time.Sleep(1 * time.Second)
	}

	for _, resp := range responses {
		log.Println(resp.StatusCode)
		log.Println(resp.Body)
		log.Println(resp.ReceivedAt)
	}

}

func getWebsiteResponse(url string) (WebsiteResponse, error) {

	// http request get to google
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return WebsiteResponse{}, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read request body: ", err)
	}

	websiteResponse := WebsiteResponse{
		StatusCode: resp.StatusCode,
		Body: string(respBody),
		ReceivedAt: time.Now(),
	}

	return websiteResponse, nil
}
	



