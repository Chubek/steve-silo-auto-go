package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MeasurementData struct {
	Success bool `json:"success"`
	Model   []struct {
		TankName          string `json:"TankName"`
		LocationName      string `json:"LocationName"`
		HiddenAlert       bool   `json:"HiddenAlert"`
		State             string `json:"State"`
		StateCode         string `json:"StateCode"`
		ImagePath         string `json:"ImagePath"`
		Percent           string `json:"Percent"`
		LatestReadingDate string `json:"LatestReadingDate"`
		Measurement       string `json:"Measurement"`
		Definitions       []struct {
			Type              string `json:"Type"`
			State             string `json:"State"`
			StateCode         string `json:"StateCode"`
			LatestReadingDate string `json:"LatestReadingDate"`
			Measurement       string `json:"Measurement"`
		} `json:"Definitions"`
	} `json:"model"`
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	loginString := "Username=sschmidt&Password=3Hc82ExR%25%40K7nO%25m&RememberMe=true&RememberMe=false&ConfirmationToken="

	requestLogin, err := http.NewRequest("POST", "https://www.mybinview.com/", strings.NewReader(loginString))
	if err != nil {
		log.Fatal(err)
	}

	requestLogin.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requestLogin.Header.Add("Content-Length", "101")
	requestLogin.Header.Add("Cache-Control", "max-age=0")
	requestLogin.Header.Add("Origin", "https://www.mybinview.com")
	requestLogin.Header.Add("Referer", "https://www.mybinview.com/")
	requestLogin.Header.Add("Host", "www.mybinview.com")
	requestLogin.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestLogin.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestLogin.Header.Add("sec-ch-ua-mobile", "?0")
	requestLogin.Header.Add("Sec-Fetch-Dest", "document")
	requestLogin.Header.Add("Sec-Fetch-Mode", "navigate")
	requestLogin.Header.Add("Sec-Fetch-Site", "same-origin")
	requestLogin.Header.Add("Sec-Fetch-User", "?1")
	requestLogin.Header.Add("Upgrade-Insecure-Requests", "1")
	requestLogin.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	requestLogin.Header.Add("Accept", "*/*")
	requestLogin.Header.Add("Accept-Encoding", "deflate")
	requestLogin.Header.Add("Connection", "keep-alive")
	requestLogin.Header.Add("Cookie", "_ga=GA1.2.832524280.1626461218; _gid=GA1.2.244243936.1626461218")

	tanksJsonVar := "{\"userID\":\"1444\"}"

	requestTanks, err := http.NewRequest("POST", "https://www.mybinview.com/Data/RefreshTanks", strings.NewReader(tanksJsonVar))
	if err != nil {
		log.Fatal(err)
	}

	requestTanks.Header.Add("Content-Type", "application/json")
	requestTanks.Header.Add("Content-Length", "17 ")
	requestTanks.Header.Add("Cache-Control", "max-age=0")
	requestTanks.Header.Add("Origin", "https://www.mybinview.com")
	requestTanks.Header.Add("Referer", "https://www.mybinview.com/Tanks")
	requestTanks.Header.Add("Host", "www.mybinview.com")
	requestTanks.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestTanks.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestTanks.Header.Add("sec-ch-ua-mobile", "?0")
	requestTanks.Header.Add("Sec-Fetch-Dest", "document")
	requestTanks.Header.Add("Sec-Fetch-Mode", "cors")
	requestTanks.Header.Add("Sec-Fetch-Site", "same-origin")
	requestTanks.Header.Add("Sec-Fetch-User", "?1")
	requestTanks.Header.Add("Upgrade-Insecure-Requests", "1")
	requestTanks.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	requestTanks.Header.Add("Accept", "*/*")
	requestTanks.Header.Add("Accept-Encoding", "deflate")
	requestTanks.Header.Add("Connection", "keep-alive")

	responseLogin, err := client.Do(requestLogin)
	if err != nil {
		log.Fatal(err)
	}

	cookiesLogin := responseLogin.Cookies()
	requestLogin.Header.Add("Cookie",
		fmt.Sprintf("_ga=GA1.2.832524280.1626461218; _gid=GA1.2.244243936.1626461218, %s=%s, %s=%s",
			cookiesLogin[0].Name, cookiesLogin[0].Value, cookiesLogin[1].Name, cookiesLogin[1].Value))

	responseTanks, err := client.Do(requestTanks)
	if err != nil {
		log.Fatal(err)
	}

	bodyRes, err := ioutil.ReadAll(responseTanks.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	var result MeasurementData

	json.Unmarshal([]byte(string(bodyRes)), &result)

	fmt.Println(result)

}
