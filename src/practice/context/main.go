package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeOutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*10000)
	defer cancel()

	// similar to clear timeout in js

	// create HTTP request

	req, err := http.NewRequestWithContext(timeOutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// get data from HTTP response
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image size of %d\n", len(imageData))

}
