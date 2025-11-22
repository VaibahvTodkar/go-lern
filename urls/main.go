package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.google.com"

func main() {
	fmt.Println("welcome to url handiling")
	fmt.Println("URL is:", myurl)

	// parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw Query:", result.RawQuery)


	querParams := result.Query()
	fmt.Println("Query Params:", querParams)
	fmt.Println("Length of Query Params:", len(querParams))
	for key, val := range querParams {
		fmt.Println("Key:", key, "Value:", val)
	}

}
