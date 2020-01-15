package server

import (
	"io/ioutil"
	"fmt"	
	"net/http"
	"os"
)

// URL contains the url to be connected or redirected
const URL = "http://kaggle.com/"

func redirectHandler(w http.ResponseWriter, r *http.Request)  {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)	
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", URL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Printf("%s", b)
}