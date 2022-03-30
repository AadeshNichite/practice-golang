package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mg sync.Mutex

func sendRequest(url string) {

	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	mg.Lock()
	fmt.Println(res.StatusCode, url)
	defer mg.Unlock()

}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("use properly")
	}

	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url)
		wg.Add(1)
	}

	wg.Wait()
}
