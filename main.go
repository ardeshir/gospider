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


