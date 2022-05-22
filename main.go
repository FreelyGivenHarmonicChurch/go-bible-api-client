package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("go-bible-api-client")

	v := viper.New()
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %w \n", err))
	}
	v.AutomaticEnv()

	client := http.Client{Timeout: time.Duration(3) * time.Second}

	url := v.GetString("bible_api.url") + "/bibles/" + v.GetString("bible_api.kjv_id")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	req.Header.Add("api-key", v.GetString("bible_api_key"))
	req.Header.Add("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

// resp, err := http.Get("http://example.com/")
// ...
// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
// ...
// resp, err := http.PostForm("http://example.com/form",
// 	url.Values{"key": {"Value"}, "id": {"123"}})

// 	resp, err := http.Get("http://example.com/")
// 	if err != nil {
// 		// handle error
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)

// 	client := &http.Client{
// 		CheckRedirect: redirectPolicyFunc,
// 	}

// 	resp, err := client.Get("http://example.com")
// 	// ...

// 	req, err := http.NewRequest("GET", "http://example.com", nil)
// 	// ...
// 	req.Header.Add("If-None-Match", `W/"wyzzy"`)
// 	resp, err := client.Do(req)

//     c := http.Client{Timeout: time.Duration(3) * time.Second}

//     req, err := http.NewRequest("GET", "http://webcode.me/ua.php", nil)

//     if err != nil {
//         log.Fatal(err)
//     }

//     req.Header.Add("User-Agent", "Go program")
//     resp, err := c.Do(req)

//     if err != nil {
//         log.Fatal(err)
//     }

//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)

//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println(string(body))
