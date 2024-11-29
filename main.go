package main

import (
	"fmt"
	"strconv"

	internal "github.com/BharaniJ27/Web-Crawler/internal/handler"
)

func main() {
	fmt.Println("Hello, Welcome to web crawler")
	var url string
	var depth string

	fmt.Println("Please enter a valid URL")
	fmt.Scan(&url)

	fmt.Println("Please enter the depth to crawl:")
	fmt.Scan(&depth)

	var depthNumber, _ = strconv.Atoi(depth)
	internal.Crawl(url, depthNumber)
}
