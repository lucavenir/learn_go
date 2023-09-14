package main

import (
	"fmt"
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

var totalErrors = 0

func main() {
	scamCh := make(chan Scam, 1_000_000_000)

	requests := 0
	completed := 0

	for i := 0; i < 10; i++ {
		log.Println("Ciao!", i)
		go scamThisScammer(scamCh)
		time.Sleep(1 * time.Nanosecond)
	}

	for r := range scamCh {
		requests += r.Attempted
		completed += r.Completed
		log.Println("Attempted: ", requests, "Completed:", completed, "Errors until now:", totalErrors)
	}
}

func scamThisScammer(ch chan Scam) {
	const getUrl = "https://launchers.world/1f49e875da6d721297ff262b54e14588"
	const anotherGetUrl = "http://clarck.online/cl/370045_smd/130/361731/2870/52"

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	for true {
		ch <- Scam{1, 0}
		sendGetRequest("GET AWS", getUrl)
		sendGetRequest("GET clark.online", anotherGetUrl)
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
		// 	log.Println("Woops, Status Code ", response.StatusCode)
		// 	body, err := io.ReadAll(response.Body)
		// 	if err != nil {
		// 		log.Println(err)
		// 		panic("Can't read the response from GET")
		// 	}
		// 	log.Println(string(body))
		totalErrors += 1
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
		// 	log.Println("Woops, Status Code ", response.StatusCode)
		// 	body, err := io.ReadAll(response.Body)
		// 	if err != nil {
		// 		log.Println(err)
		// 		panic("Can't read the response from GET")
		// 	}
		// 	log.Println(string(body))
		totalErrors += 1
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
