package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// some global variables

var applicationStaus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalURLCount int
var wg sync.WaitGroup

var v1 int

func readURLs(statusChannel chan int, textChannel chan string) {
	
	time.Sleep(time.Millisecond * 1)
	fmt.Println("Grapping...", len(urls), "urls\n")
	for i := 0, i < totalURLCount; i++ {
		fmt.Println("Url: ", i, ulrs[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
	
		textChannel <- string(text)

		if err != nil {
			fmt.Println("No HTML Body")
		}

		statusChannel <- 0
	}
}

func addToScrapedText(textChannel chan string, processChannel chan bool){
	for {
		select {
		case pC := <- processChannel:
			if pC == true {
			 // hang on
			}
		if pC == false {
			close(textChannel)
			close(processChannel)
		}
		case tC := <- textChannel:
		fullText += tC
 	      }
	}
}
	
