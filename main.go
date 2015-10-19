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

		// fmt.Printf("Url: %d, %s\n", i, urls[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
                fmt.Printf("Resp Code: %d ->", resp.StatusCode)
		textChannel <- string(text)
	        fmt.Printf("Url: %d, %s\n", i, urls[i])	
		if err != nil {
			fmt.Println("No HTML Body")
		}
		
		statusChannel <- 0
	}
}

func addToScrapedText(statusChannel chan int, textChannel chan string, processChannel chan bool){
	
	for {
		select {
		case pC := <- processChannel:
			if pC == true {
			 // hang on
			}
		if pC == false {
			close(textChannel)
			close(processChannel)
                 break
		}
		case tC := <- textChannel:
		fullText += tC
                urlsProcessed++
                fmt.Printf("UrlsProcessed: %d in tC adding text\n", urlsProcessed)
                if(urlsProcessed  == totalURLCount) {
                    applicationStatus = false
                } 
 	      }
	}
}

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool){

	for {
	    select {
		case status := <- statusChannel:
		fmt.Println("urlProc:", urlsProcessed, " TotalUrls:", totalURLCount)
		// urlsProcessed++
		if status == 0 {
			fmt.Printf("evalStatus:  %d, urlsProcessed: %d\n", status, urlsProcessed)
		}
		if status == 1 {
                        fmt.Printf("evalStatus:  %d\n", status)
			close(statusChannel)
		}
       	   }
             /* if urlsProcessed == totalURLCount {
		fmt.Println("Read all top-level URLs")
		processChannel <- false
		applicationStatus = false
             break
             } */
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
	
	// go evaluateStatus(statusChannel, textChannel, processChannel)

	go readURLs(statusChannel, textChannel)
	
	go addToScrapedText(statusChannel, textChannel, processChannel)

	go evaluateStatus(statusChannel, textChannel, processChannel)

       for {
	    select {
         	case sC := <- statusChannel:
	               fmt.Printf("StatusChannel %d, Status: %t\n", sC,applicationStatus)
	    } 

             if applicationStatus == false {
		 fmt.Println("fullText...")
		// fmt.Println("Done, bye!")
		break
	    }
	}
    fmt.Println("End of program!")
}		
