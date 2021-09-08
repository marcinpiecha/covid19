package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/marcinpiecha/covid19/config"
	"github.com/marcinpiecha/covid19/stats"
)

type M map[string]interface{}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	//fmt.Println("server is running...")

	//1st request
	url := "https://covid19.mathdro.id/api"
	//url := "http://api.open-notify.org/astros.json"

	coronaClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "coronavirus-stats")

	res, getErr := coronaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data := stats.Corona{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//end of first Request

	url2 := "https://covid19.mathdro.id/api/confirmed"

	coronaClient2 := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	req2, err2 := http.NewRequest(http.MethodGet, url2, nil)
	if err2 != nil {
		log.Fatal(err2)
	}

	res2, getErr2 := coronaClient2.Do(req2)
	if getErr != nil {
		log.Fatal(getErr2)
	}

	body2, readErr := ioutil.ReadAll(res2.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data2 := stats.Global{}
	jsonErr2 := json.Unmarshal(body2, &data2)
	if jsonErr != nil {
		log.Fatal(jsonErr2)
	}

	//fmt.Println(data.Confirmed.Value)

	config.TPL.ExecuteTemplate(w, "index.gohtml", M{
		"data":  data,
		"data2": data2,
	})
}

func Individual(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	//fmt.Println("server is running...")

	myurl := r.URL.String()

	link := stats.Url{Country: myurl[1:]}

	fmt.Println(link)
	//fmt.Printf("%v", myurl)
	//1st request
	url := "https://covid19.mathdro.id/api/countries" + myurl
	//url := "http://api.open-notify.org/astros.json"

	coronaClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "coronavirus-stats")

	res, getErr := coronaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data := stats.Country{}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//end of first Request

	//fmt.Println(data.Confirmed.Value)

	config.TPL.ExecuteTemplate(w, "individual.gohtml", M{
		"data": data,
		"link": link,
	})
}
