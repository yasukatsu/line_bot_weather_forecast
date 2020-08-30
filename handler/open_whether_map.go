package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// StructOpenWhetherMap ...
type StructOpenWhetherMap struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

var url = "https://community-open-weather-map.p.rapidapi.com/weather?q=Tokyo,jp"

// OpenWhetherMap ...
func OpenWhetherMap(w http.ResponseWriter, r *http.Request) {

	res, err := openWhetherMapResponse(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(res)
}

// NowTemp ...
func NowTemp(w http.ResponseWriter, r *http.Request) {

	result := nowTemp()
	if result == "" {
		w.Write([]byte("エラーだお"))
		return
	}

	w.Write([]byte(result))
}

func nowTemp() string {
	res, err := openWhetherMapResponse(url)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	s := convertToStruct(res)
	if s == nil {
		log.Fatalf("fatal convertToStruct")
		return ""
	}

	// ケルビン(K) から摂氏(℃) に変換するに、-273.15 をする。
	temp := s.Main.Temp - float64(273.15)
	v := fmt.Sprintf("今の東京の温度は%v度です。\n", int(temp))

	return v
}

func convertToStruct(value []byte) *StructOpenWhetherMap {
	s := new(StructOpenWhetherMap)

	err := json.Unmarshal(value, s)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return s
}

func openWhetherMapResponse(url string) ([]byte, error) {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", os.Getenv("X_RAPIDAPI_HOST"))
	req.Header.Add("x-rapidapi-key", os.Getenv("X_RAPIDAPI_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
