package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=Tokyo,jp"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", os.Getenv("X_RAPIDAPI_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("X_RAPIDAPI_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))

}
