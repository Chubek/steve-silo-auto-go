package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tebeka/selenium"
)

func scrape_m2web() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).

	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "vendor/geckodriver"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox", "headless": true}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://operator:operator123@us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm"); err != nil {
		panic(err)
	}
	file, err := os.Open("sc.png") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data, err := wd.Screenshot()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	scrape_m2web()
}