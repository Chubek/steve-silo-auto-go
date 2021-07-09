package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/PuerkitoBio/goquery"
)

func scrape_m2web() {
	client := &http.Client{}

	responseInitial, err := http.Get("https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm")
	if err != nil {
		log.Fatal(err)
	}
	defer responseInitial.Body.Close()

	dataRequestFirst := url.Values{
		"account":  {"valleycarriers"},
		"username": {"operator"},
		"password": {"operator123"},
		"connect":  {"connect"},
		"attempt":  {"0"},
	}

	requestFirst, err := http.NewRequest("POST", "https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm", strings.NewReader(dataRequestFirst.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	requestFirst.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	requestFirst.Header.Add("Accept-Encoding", "gzip, deflate, br")
	requestFirst.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestFirst.Header.Add("Cache-Control", "max-age=0")
	requestFirst.Header.Add("Connection", "keep-alive")
	requestFirst.Header.Add("Content-Length", strconv.Itoa(len(dataRequestFirst.Encode())))
	requestFirst.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rInCookies := responseInitial.Cookies()
	requestFirst.Header.Add("Cookie", fmt.Sprintf("%s=%s; %s=%s", rInCookies[0].Name, rInCookies[0].Value, rInCookies[1].Name, rInCookies[1].Value))
	requestFirst.Header.Add("Host", "us2.m2web.talk2m.com")
	requestFirst.Header.Add("Origin", "https://us2.m2web.talk2m.com")
	requestFirst.Header.Add("Referer", "https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm")
	requestFirst.Header.Add("sec-ch-ua", "Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestFirst.Header.Add("sec-ch-ua-mobile", "?0")
	requestFirst.Header.Add("Sec-Fetch-Dest", "document")
	requestFirst.Header.Add("Sec-Fetch-Site", "same-origin")
	requestFirst.Header.Add("Sec-Fetch-Mode", "navigate")
	requestFirst.Header.Add("Sec-Fetch-User", "?1")
	requestFirst.Header.Add("Upgrade-Insecure-Requests", "1")
	requestFirst.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.3")

	requestValidate, err := http.NewRequest("GET", "https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm", strings.NewReader(dataRequestSecond.Encode()))

	dataRequestSecond := url.Values{}
	dataRequestSecond.Set("_vars", "Silo_3_ROC,Silo_1_Level,Silo_2_Level,Silo_3_Level,$script4$viewon!=time$$/script4$,Silo_2_ROC,Silo_1_ROC")

	requestSecond, err := http.NewRequest("POST", "https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/rcgi.bin/vows/readVars", strings.NewReader(dataRequestSecond.Encode()))

	if err != nil {
		log.Fatal(err)
	}
	requestSecond.Header.Add("Accept", "application/xml, text/xml, */*; q=0.01")
	requestSecond.Header.Add("Accept-Encoding", "gzip, deflate, br")
	requestSecond.Header.Add("Accept-Language", "en-US,en;q=0.9")
	requestSecond.Header.Add("Authorization", "Basic b3BlcmF0b3I6b3BlcmF0b3IxMjM=")
	requestSecond.Header.Add("cache-Control", " no-cache")
	requestSecond.Header.Add("Connection", "keep-alive")
	requestSecond.Header.Add("Content-Length", strconv.Itoa(len(dataRequestSecond.Encode())))
	requestSecond.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rTwoInCookies := responseFirst.Cookies()
	requestSecond.Header.Add("Cookie", fmt.Sprintf("m2web.token=%s; m2webskin=%s; m2websession=%s", rTwoInCookies[1].Value, rTwoInCookies[0].Value, rTwoInCookies[1].Value))
	requestSecond.Header.Add("Host", "us2.m2web.talk2m.com")
	requestSecond.Header.Add("Origin", "https://us2.m2web.talk2m.com")
	requestSecond.Header.Add("Referer", "https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm")
	requestSecond.Header.Add("sec-ch-ua", "Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	requestSecond.Header.Add("sec-ch-ua-mobile", "?0")
	requestSecond.Header.Add("Sec-Fetch-Dest", "empty")
	requestSecond.Header.Add("Sec-Fetch-Mode", "cors")
	requestSecond.Header.Add("Sec-Fetch-Site", "same-origin")
	requestSecond.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.3")
	requestSecond.Header.Add("X-Requested-With", "XMLHttpRequest")

	responseFinal, err := client.Do(requestSecond)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(responseFinal.Status)
	defer responseFinal.Body.Close()
	body, err := ioutil.ReadAll(responseFinal.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}

func main() {
	scrape_m2web()
}
