package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// OpenWhetherMap ...
func OpenWhetherMap(w http.ResponseWriter, r *http.Request) {

	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=Tokyo,jp"

	res, err := openWhetherMapResponse(url)
	if err != nil {
		return
	}

	w.Write(res)
}

func openWhetherMapResponse(url string) ([]byte, error) {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", os.Getenv("X_RAPIDAPI_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("X_RAPIDAPI_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
