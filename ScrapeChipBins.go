package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anthonynsimon/bild/effect"
	"github.com/disintegration/imaging"
	"github.com/lazureykis/dotenv"
	"github.com/mxschmitt/playwright-go"
	"github.com/oliamb/cutter"
)

type ChipBinsOCRRes struct {
	ParsedResults []struct {
		TextOverlay struct {
			Lines []struct {
				LineText string `json:"LineText"`
				Words    []struct {
					WordText string  `json:"WordText"`
					Left     float64 `json:"Left"`
					Top      float64 `json:"Top"`
					Height   float64 `json:"Height"`
					Width    float64 `json:"Width"`
				} `json:"Words"`
				MaxHeight float64 `json:"MaxHeight"`
				MinTop    float64 `json:"MinTop"`
			} `json:"Lines"`
			HasOverlay bool   `json:"HasOverlay"`
			Message    string `json:"Message"`
		} `json:"TextOverlay"`
		TextOrientation   string `json:"TextOrientation"`
		FileParseExitCode int    `json:"FileParseExitCode"`
		ParsedText        string `json:"ParsedText"`
		ErrorMessage      string `json:"ErrorMessage"`
		ErrorDetails      string `json:"ErrorDetails"`
	} `json:"ParsedResults"`
	OCRExitCode                  int    `json:"OCRExitCode"`
	IsErroredOnProcessing        bool   `json:"IsErroredOnProcessing"`
	ProcessingTimeInMilliseconds string `json:"ProcessingTimeInMilliseconds"`
}

func ScrapeChipBins() (x ChipBinsOCRRes) {
	dotenv.Go()
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("http://chipbins.orovillereload.com/rvscale.htm?session=null&info=/assets/json/wxga-full-x1-16.json", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	time.Sleep(5 * time.Second)

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	existingImageFile, err := os.Open("foo.png")
	if err != nil {
		log.Fatal(err)
	}

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		log.Fatal(err)
	}

	croppedImg, err := cutter.Crop(loadedImage, cutter.Config{
		Width:  loadedImage.Bounds().Dx() - 120,
		Anchor: image.Point{100, 100},
		Height: 150,
	})
	if err != nil {
		log.Fatal(err)
	}

	brtImage := imaging.AdjustBrightness(croppedImg, 40)
	cnstImage := imaging.AdjustContrast(brtImage, 140)
	sigImage := imaging.AdjustSigmoid(cnstImage, 0.5, 0.5)
	gmImage := imaging.AdjustGamma(sigImage, 0.8)
	satImage := imaging.Sharpen(gmImage, 0.5)
	finImage := imaging.AdjustContrast(satImage, 240)
	erodedImage := effect.Erode(finImage, 1)

	imgRezied := imaging.Resize(erodedImage, croppedImg.Bounds().Dx()*4, croppedImg.Bounds().Dy()*4, imaging.Lanczos)

	buf := new(bytes.Buffer)
	png.Encode(buf, imgRezied)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	b64WithSignature := fmt.Sprintf("data:image/png;base64,%s", imgBase64Str)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	dataOcr := url.Values{}
	dataOcr.Set("scale", "true")
	dataOcr.Set("isTable", "true")
	dataOcr.Set("filetype", "PNG")
	dataOcr.Set("isOverlayRequired", "false")
	dataOcr.Set("OCREngine", "2")
	dataOcr.Set("base64Image", b64WithSignature)

	requestOCR, err := http.NewRequest("POST", "https://api.ocr.space/Parse/Image", strings.NewReader(dataOcr.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	requestOCR.Header.Add("apikey", os.Getenv("API_KEY"))
	requestOCR.Header.Add("Accept", "*/*")
	requestOCR.Header.Add("Connection", "keep-alive")
	requestOCR.Header.Add("Accept-Encoding", "deflate")
	requestOCR.Header.Add("User-Agent", "Omar's Coming Yo!")
	requestOCR.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requestOCR.Header.Add("Content-Length", strconv.Itoa(len(dataOcr.Encode())))

	res, err := client.Do(requestOCR)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	var ocrResults ChipBinsOCRRes

	json.Unmarshal(body, &ocrResults)

	fmt.Printf("%+v\n", ocrResults)

	return ocrResults
}
