package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type picdata struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const baseurl = "https://xkcd.com/"
const tag = "/info.0.json"
const failTag = 4

func FetchingDataGOPL() {
	fullData := []picdata{}
	checkFailTag := 0
	shouldVisitNext := true

	for i := 1; shouldVisitNext; i++ {

		num := strconv.Itoa(i)

		resp, err := http.Get(baseurl + num + tag)
		if err != nil {
			fmt.Println("Hellyeahbitch")
			checkFailTag++

			fmt.Fprintln(os.Stderr, err)
		}

		defer resp.Body.Close()
		var body picdata
		fmt.Println("body",resp.Body)
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			checkFailTag++

			fmt.Fprintln(os.Stderr, err)
		}
		fullData = append(fullData, body)
		// fmt.Println(body)
		if checkFailTag > failTag {
			shouldVisitNext = false
		}

	}

	//Saving it to an external file
	file, _ := json.Marshal(fullData)

	_ = os.WriteFile("test.json", file, 0644)

}
