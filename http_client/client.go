package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < 10000; i++ {
		time.Sleep(100 * time.Millisecond)
		wg.Add(1)
		go scamChain()
	}
	wg.Wait()
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

func scamChain() {
	const first = "https://ui-servizi-infonet.162-254-35-167.cprapid.com/std-lsc/it/login2.php"
	const second = "https://ui-servizi-infonet.162-254-35-167.cprapid.com/std-lsc/it/login3.php"
	const third = "https://ui-servizi-infonet.162-254-35-167.cprapid.com/std-lsc/it/login4.php"

	defer wg.Done()
	for true {
		time.Sleep(100 * time.Millisecond)

		defer func() {
			if r := recover(); r != nil {
				log.Print("Recovered from panic: ")
				log.Println(r)
			}
		}()
		scamTheScammer("First request", first, map[string]any{
			"cod":   stringWithCharset(7),
			"pin":   stringWithCharset(5),
			"Invia": "",
		})
		scamTheScammer("Second request", second, map[string]any{
			"otp":   stringWithCharset(4),
			"Invia": "",
		})
		scamTheScammer("Third request", third, map[string]any{
			"otp":   stringWithCharset(4),
			"Invia": "",
		})
	}
}

func scamTheScammer(title string, url string, data map[string]any) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		panic("Can't convert data to json")
	}
	reader := bytes.NewReader(j)
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		log.Println(err)
		panic("Can't actually send this request")
	}
	defer response.Body.Close()
	log.Printf("[%s] Sent: ", title)
	log.Println(data)
	log.Printf("[%s] Received: ", title)
	log.Println(response.StatusCode)
}
