package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var items []interface{}

var item map[string]interface{}

func main() {
	numberOfJobs := flag.Int("n", 20, "number of jobs you want to retrieve")
	flag.Parse()

	hnURL := "https://hacker-news.firebaseio.com//v0/jobstories.json"

	resp, err := http.Get(hnURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get job items from Hacker news link  because of error %v", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get read data from Hacker news link %v because of error %v", hnURL, err)
		os.Exit(1)
	}

	err2 := json.Unmarshal(body, &items)

	if err2 != nil {
		fmt.Printf("Error parsing news item %v", err2)
	}
	resp.Body.Close()

	for _, itemValue := range items[:*numberOfJobs] {
		itemNumber := itemValue.(float64)
		itemInt := int(itemNumber)

		itemURL := "https://hacker-news.firebaseio.com/v0/item/" + fmt.Sprint(itemInt) + ".json"

		getItem(itemURL)
	}

}

func getItem(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	println("status is", resp.StatusCode)
	er := json.Unmarshal(b, &item)

	if er != nil {
		fmt.Printf("Error in getting news item is %v", er)
	}

	fmt.Printf("%v \n URL -> %v", item["title"], item["url"])
}
