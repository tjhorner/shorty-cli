package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type shortenRes struct {
	Error  *string `json:"error"`
	Result *struct {
		Suffix string `json:"suffix"`
		URL    string `json:"url"`
	} `json:"result"`
}

func oops(err interface{}) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	base := os.Getenv("SHORTY_BASE_URL")
	if base == "" {
		oops("no base URL provided. put it in the SHORTY_BASE_URL env variable. (e.g., https://example.com)")
	}

	flag.Parse()

	longURL := flag.Arg(0)
	if longURL == "" {
		oops("no URL provided. usage: shorty <url> [suffix]")
	}

	suffix := flag.Arg(1)

	form := url.Values{}
	form.Add("url", longURL)
	form.Add("suffix", suffix)

	req, err := http.NewRequest("POST", base+"/api/shorten", strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		oops(err)
	}
	defer res.Body.Close()

	var resp shortenRes

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&resp)
	if err != nil {
		oops(err)
	}

	if resp.Error != nil {
		oops(*resp.Error)
	}

	fmt.Printf(base+"/%s\n", resp.Result.Suffix)
}
