package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	_"golang.org/x/net/html/charset"
	_"bytes"
)


type Server struct {
	XMLName xml.Name `xml:"server"`
	Text    string   `xml:",chardata"`
	Vars    struct {
		Text string `xml:",chardata"`
		Var  []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Value   string `xml:"value,attr"`
			Quality string `xml:"quality,attr"`
		} `xml:"var"`
	} `xml:"vars"`

}

func scrape_m2web() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	requestInit, err := http.NewRequest("GET", "https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm", nil)
	if err != nil {
		log.Fatal(err)
	}

	requestInit.Header.Add("Host", "us2.m2web.talk2m.com")
	requestInit.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestInit.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestInit.Header.Add("sec-ch-ua-mobile", "?0")
	requestInit.Header.Add("Sec-Fetch-Dest", "document")
	requestInit.Header.Add("Sec-Fetch-Mode", "navigate")
	requestInit.Header.Add("Sec-Fetch-Site", "none")
	requestInit.Header.Add("Sec-Fetch-User", "?1")
	requestInit.Header.Add("Upgrade-Insecure-Requests", "1")
	requestInit.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	requestInit.Header.Add("Accept", "*/*")
	requestInit.Header.Add("Accept-Encoding", "gzip, deflate, br")
	requestInit.Header.Add("Connection", "keep-alive")
	requestInit.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3IxMjM=")

	loginString := "account=valleycarriers&username=operator&password=operator123&connect=connect&attempt=0"

	requestLogin, err := http.NewRequest("POST", "https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm", strings.NewReader(loginString))
	if err != nil {
		log.Fatal(err)
	}

	requestLogin.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requestLogin.Header.Add("Content-Length", "87")
	requestLogin.Header.Add("Cache-Control", "max-age=0")
	requestLogin.Header.Add("Origin", "https://us2.m2web.talk2m.com")
	requestLogin.Header.Add("Referer", "https://us2.m2web.talk2m.com/valleycarriers/Gorman Bros/usr/viewon/Overview.shtm")
	requestLogin.Header.Add("Host", "us2.m2web.talk2m.com")
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
	requestLogin.Header.Add("Accept-Encoding", "gzip, deflate, br")
	requestLogin.Header.Add("Connection", "keep-alive")
	requestLogin.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3IxMjM=")

	requestValidate, err := http.NewRequest("GET", "https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm", nil)
	if err != nil {
		log.Fatal(err)
	}

	requestValidate.Header.Add("Host", "us2.m2web.talk2m.com")
	requestValidate.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestValidate.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestValidate.Header.Add("sec-ch-ua-mobile", "?0")
	requestValidate.Header.Add("Sec-Fetch-Dest", "document")
	requestValidate.Header.Add("Sec-Fetch-Mode", "navigate")
	requestValidate.Header.Add("Sec-Fetch-Site", "none")
	requestValidate.Header.Add("Sec-Fetch-User", "?1")
	requestValidate.Header.Add("Upgrade-Insecure-Requests", "1")
	requestValidate.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	requestValidate.Header.Add("Accept", "*/*")
	requestValidate.Header.Add("Accept-Encoding", "gzip, deflate, br")
	requestValidate.Header.Add("Connection", "keep-alive")
	requestValidate.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3IxMjM=")
	requestValidate.Header.Add("Cache-Control", "max-age=0")
	requestValidate.Header.Add("Referer", "https://us2.m2web.talk2m.com/valleycarriers/Gorman Bros/usr/viewon/Overview.shtm")

	varsString := "_vars=Silo_3_ROC%2CSilo_1_Level%2CSilo_2_Level%2CSilo_3_Level%2C%24script4%24viewon!%3Dtime%24%24%2Fscript4%24%2CSilo_2_ROC%2CSilo_1_ROC"

	requestReadVars, err := http.NewRequest("POST", "https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman Bros/rcgi.bin/vows/readVars", strings.NewReader(varsString))
	if err != nil {
		log.Fatal(err)
	}

	requestReadVars.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	requestReadVars.Header.Add("Content-Length", "136")
	requestReadVars.Header.Add("Cache-Control", "no-cache")
	requestReadVars.Header.Add("Origin", "https://us2.m2web.talk2m.com")
	requestReadVars.Header.Add("Referer", "https://us2.m2web.talk2m.com/valleycarriers/Gorman Bros/usr/viewon/Overview.shtm")
	requestReadVars.Header.Add("Host", "us2.m2web.talk2m.com")
	requestReadVars.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestReadVars.Header.Add("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestReadVars.Header.Add("sec-ch-ua-mobile", "?0")
	requestReadVars.Header.Add("Sec-Fetch-Dest", "empty")
	requestReadVars.Header.Add("Sec-Fetch-Mode", "cors")
	requestReadVars.Header.Add("Sec-Fetch-Site", "same-origin")
	requestReadVars.Header.Add("Upgrade-Insecure-Requests", "1")
	requestReadVars.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	requestReadVars.Header.Add("Accept", "*/*")
	requestReadVars.Header.Add("Accept-Encoding", "deflate")
	requestReadVars.Header.Add("Connection", "keep-alive")
	requestReadVars.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3IxMjM=")
	requestReadVars.Header.Add("X-Requested-With", "XMLHttpRequest")

	responseInit, err := client.Do(requestInit)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response Init Header", responseInit.Header)

	bodyResInit, err := ioutil.ReadAll(responseInit.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	log.Println("Response Init Body", string(bodyResInit))

	fmt.Println("---------------------------------------------------------")

	for _, cookie := range responseInit.Cookies() {
		requestLogin.AddCookie(cookie)
	}

	responseLogin, err := client.Do(requestLogin)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response Login Header", responseLogin.Header)

	bodyResLogin, err := ioutil.ReadAll(responseLogin.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	log.Println("Response Login Body", string(bodyResLogin))

	fmt.Println("---------------------------------------------------------")


	for _, cookie := range responseLogin.Cookies() {
		requestValidate.AddCookie(cookie)
	}

	responseValidate, err := client.Do(requestValidate)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response Validate Header", responseValidate.Header)

	fmt.Println("---------------------------------------------------------")


	for _, cookie := range responseValidate.Cookies() {
		requestReadVars.AddCookie(cookie)
	}

	for _, cookie := range responseLogin.Cookies() {
		requestReadVars.AddCookie(cookie)
	}

	log.Println("Request ReadVars Header", requestReadVars.Header)

	responseReadVars, err := client.Do(requestReadVars)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response ReadVars Header", responseReadVars.Header)

	
	bodyRes, err := ioutil.ReadAll(responseReadVars.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	err2 := ioutil.WriteFile("dat2", bodyRes, 0644)
    if err2 != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	/*
	reader := bytes.NewReader(bodyRes)
	decoder := xml.NewDecoder(reader)

	decoder.CharsetReader = charset.NewReaderLabel
	

	var parsed Server

	err = decoder.Decode(&parsed)
	*/
	fmt.Println(fmt.Sprintf("%s", string(bodyRes)))


}

func main() {
	scrape_m2web()
}
