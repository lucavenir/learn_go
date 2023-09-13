package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Scam struct {
	Attempted int
	Completed int
}

func main() {
	scamCh := make(chan Scam, 1_000_000)

	requests := 0
	completed := 0

	for i := 0; i < 15_000; i++ {
		log.Println("Ciao!", i)
		go scamThisScammer(scamCh)
	}

	for r := range scamCh {
		requests += r.Attempted
		completed += r.Completed
		log.Println("Attempted: ", requests, "Completed:", completed)
	}
}

func scamThisScammer(ch chan Scam) {
	const getUrl = "https://launcherpod.info/5f39888832e18d47260016cebf177f14"
	const postUrl = "https://amplinesrv.com/survey/saveAnswer"

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	for true {
		ch <- Scam{1, 0}
		sendGetRequest("GET launcherpod.info", getUrl)
		sendPostRequest("POST amplinesrv.com", postUrl)
		ch <- Scam{0, 1}
	}
}

func sendPostRequest(title string, link string) {
	values := url.Values{
		"bid":  {stringWithCharset(9)},
		"sid":  {stringWithCharset(2)},
		"lid":  {stringWithCharset(1)},
		"cnt":  {stringWithCharset(3)},
		"qid":  {stringWithCharset(3)},
		"aid":  {stringWithCharset(3)},
		"step": {stringWithCharset(1)},
		"pos":  {fmt.Sprintf("%t", flipACoin())},
		"cmp":  {"MailSurvey"},
	}
	response, err := http.PostForm(link, values)
	if err != nil {
		log.Println(err)
		panic("Can't actually send this request")
	}
	defer response.Body.Close()
	// log.Printf("[%s] Sent: ", title)
	// log.Println(values)
	// log.Printf("[%s] Received: ", title)
	// log.Println(response.StatusCode)
	if response.StatusCode != 200 {
		log.Println("Woops, Status Code ", response.StatusCode)
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			panic("Can't read the response from GET")
		}
		log.Println(string(body))
	}
}

func sendGetRequest(title string, url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		panic("Can't send this request")
	}
	defer response.Body.Close()
	// log.Printf("[%s] Received: ", title)
	// log.Println(response.StatusCode)
	if response.StatusCode != 200 {
		log.Println("Woops, Status Code ", response.StatusCode)
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			panic("Can't read the response from GET")
		}
		log.Println(string(body))
	}
}

func flipACoin() bool {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := seededRand.Float64()
	return random < 0.5
}

func stringWithCharset(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
