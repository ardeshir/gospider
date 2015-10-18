package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// some global variables

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalURLCount int
//  var wg sync.WaitGroup
var proc int 
var v1 int

func readURLs(statusChannel chan int, textChannel chan string) {
	
	time.Sleep(time.Millisecond * 3)
	fmt.Println("Grapping...", len(urls), " urls")
	for i := 0; i < totalURLCount; i++ {
		fmt.Printf("Url: %d, %s\n", i, urls[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
                fmt.Println(resp.StatusCode)
  		/* seeing text
                if err == nil {
		fmt.Printf("%s", string(text))
                }
		*/ 
	
		textChannel <- string(text)
                urlsProcessed++
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
                        fmt.Printf("Proc %d in pC true", proc)
			}
		if pC == false {
			close(textChannel)
			close(processChannel)
                 break
		}
		case tC := <- textChannel:
		fullText += tC
                fmt.Printf("Proc %d in tC adding text\n", proc)
                proc++ 
 	      }
	}
}

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool){

	for {
	    select {
		case status := <- statusChannel:
		fmt.Println("urlProc:", urlsProcessed, " TotalUrls:", totalURLCount)
		urlsProcessed++
		if status == 0 {
			fmt.Printf("Got url: %d\n", urlsProcessed)
		}
		if status == 1 {
			close(statusChannel)
		}
		if urlsProcessed == totalURLCount {
			fmt.Println("Read all top-level URLs")
			processChannel <- false
			applicationStatus = false
                break
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
	v1 = 1
	urls = append(urls, "http://admin.hennepintech.edu")
	urls = append(urls, "http://joymonk.com")
	urls = append(urls, "http://metalearn.org")
	urls = append(urls, "http://scholarlylife.com")
	urls = append(urls, "http://univrs.io")

	fmt.Println("Starting Sprider version  ", v1)
	
	urlsProcessed = 0 
	totalURLCount = len(urls)
	
	go evaluateStatus(statusChannel, textChannel, processChannel)

	go readURLs(statusChannel, textChannel)
	
	go addToScrapedText(textChannel, processChannel)

	for {
		if applicationStatus == false {
			fmt.Println(fullText)
			fmt.Println("=-=-=-=-=-=-=-=-=")
			fmt.Println("Done, bye!")
		break
		}
	    select {
		case sC := <- statusChannel:
	fmt.Printf("StatusChannel %d, Status: %t\n", sC,applicationStatus)
             
	   }
	}
}		
