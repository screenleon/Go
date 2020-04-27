package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet() {
	resp, err := http.Get("https://tw.yahoo.com/")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("http get 1:", string(body))
}

func httpPost() {
	resp, err := http.Post("https://tw.yahoo.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("http post 1:", string(body))
}

func httpPost2() {
	resp, err := http.PostForm("https://tw.yahoo.com/",
		url.Values{"key": {"Value"}, "id": {"test"}})

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("http post 2:", string(body))
}

func httpPostWithCookis() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://tw.yahoo.com/", strings.NewReader("name=test"))
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=test")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	fmt.Println("http post with cookie:", string(body))
}

func main() {
	// httpGet()
	httpPost()
	// httpPost2()
	// httpPostWithCookis()
}
