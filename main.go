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
	fmt.Println("Grapping...", len(urls), " urls")
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

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool){

	for {
	    select {
		case status := <- statusChannel:
		fmt.Println(urlsProcessed, totalURLCount)
		urlsProcessed++
		if status == 0 {
			fmt.Println("Got url")
		}
		if status == 1 {
			close(statusChannel)
		}
		if urlsProcessed == totalURLCount {
			fmt.Println("Read all top-level URLs")
			processChannel <- false
			applicationStatus = false
		}
	   }
	}
}

// adding main 

func main(){
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)
	totalURLCount = 0

	urls = append(urls, "http://ardeshir.org")
	urls = append(urls, "http://joymonk.com")
	urls = append(urls, "http://metalearn.org")
	urls = append(urls, "http://scholarlylife.com")
	urls = append(urls, "http://univrs.io")

	fmt.Println("Starting sprider...")
	
	urlsProcessed = 0
	totalURLCount = len(urls)
	
	go evaluateStatus(statusChannel, textChannel, processChannel)

	go readURLs(statusChannel, processChannel)

	for {
		if applicationStatus == false {
			fmt.Println(fullText)
			fmt.Println("=-=-=-=-=-=-=-=-=")
			fmt.Println("Done, bye!")
		break
		}
	    select {
		case sC := <- statusChannel:
			fmt.Println("Message on StatusChannel", sC)
	   }
	}
}		
