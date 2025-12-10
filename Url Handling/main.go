package main

import (
	"fmt"
	"net/url"
)

func main() {
	Myurl := "https://www.example.com:8080/path/to/resource?query=param#section"
	fmt.Println("Original URL:", Myurl)

	parsedUrl, err := url.Parse(Myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	// Modifying the URL
	parsedUrl.Scheme = "http"
	parsedUrl.Host = "www.changedexample.com:9090"
	parsedUrl.Path = "/new/path"
	parsedUrl.RawQuery = "newquery=newparam"
	parsedUrl.Fragment = "newsection"

	// Reconstructing the modified URL
	modifiedUrl := parsedUrl.String()
	fmt.Println("Modified URL:", modifiedUrl)
}
