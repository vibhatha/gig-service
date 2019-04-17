// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/utility/requesthandlers"
	"bytes"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"io"
	"net/url"
	"os"
)

var visited = make(map[string]bool)
var apiUrl = "http://localhost:9000/api/add"

/**
	get tender url and append page number
	get page html and query media bodies
	get notice link for pdf
	download pdf
	read pdf text/image
	extract pdf content using NER/Tesseract
	save to mongo
 */

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	fmt.Println(args[0])
	//
	//for pageNo := range 1,20 {
	//	response := enqueue(uri, queue)
	//	entity := decoder.DecodeSource(response, uri)
	//	_, err := requesthandlers.PostRequest(apiUrl, entity)
	//	if err != nil {
	//		fmt.Println(err.Error(),uri)
	//	}
	//}
}

func enqueue(uri string, queue chan string) *bytes.Buffer {
	fmt.Println("fetching", uri)
	visited[uri] = true

	resp, err := requesthandlers.GetRequest(uri)

	if err != nil {
		return &bytes.Buffer{}
	}
	var bufferedResponse bytes.Buffer
	response := io.TeeReader(resp.Body, &bufferedResponse)
	links := collectlinks.All(response)
	defer resp.Body.Close()

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if uri != "" {
			if !visited[absolute] {
				go func() { queue <- absolute }()
			}
		}
	}
	return &bufferedResponse
}

func fixUrl(href, base string) (string) {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}