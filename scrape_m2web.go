package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tebeka/selenium"
)

// This example shows how to navigate to a http://play.golang.org page, input a
// short program, run it, and inspect its output.
//
// If you want to actually run this example:
//
//   1. Ensure the file paths at the top of the function are correct.
//   2. Remove the word "Example" from the comment at the bottom of the
//      function.
//   3. Run:
//      go test -test.run=Example$ github.com/tebeka/selenium
func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).

	seleniumPath := "selenium-server-standalone-3.141.59.jar"
	driverPath := "chromedriver"
	port := 9688

	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),       // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(driverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),        // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "HTMUnit"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://us2.m2web.talk2m.com/valleycarriers/Gorman%20Bros/usr/viewon/Overview.shtm"); err != nil {
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

	/*
		// Get a reference to the text box containing code.
		elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
		if err != nil {
			panic(err)
		}
		// Remove the boilerplate code already in the text box.
		if err := elem.Clear(); err != nil {
			panic(err)
		}

		// Enter some new code in text box.
		err = elem.SendKeys(`
			package main
			import "fmt"
			func main() {
				fmt.Println("Hello WebDriver!")
			}
		`)
		if err != nil {
			panic(err)
		}

		// Click the run button.
		btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
		if err != nil {
			panic(err)
		}
		if err := btn.Click(); err != nil {
			panic(err)
		}

		// Wait for the program to finish running and get the output.
		outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
		if err != nil {
			panic(err)
		}

		var output string
		for {
			output, err = outputDiv.Text()
			if err != nil {
				panic(err)
			}
			if output != "Waiting for remote server..." {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}

		fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
	*/
	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.
}
