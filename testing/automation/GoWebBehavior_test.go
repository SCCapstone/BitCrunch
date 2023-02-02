package testing

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/go-rod/rod"
)

// https://github.com/go-rod/rod/blob/master/examples_test.go

func TestMain(m *testing.M) {
	fmt.Println("Please note: in order to properly test, please read the Readme!! Otherwise, the browser will NOT connect properly!")
	fmt.Println("Building and Activating GoWeb...")
	fmt.Println("Making sure that we are in the correct directory...")
	cherry := os.Chdir("../../")
	if cherry != nil {
		// Something happened when trying to change dir!
		fmt.Println("chdir didn't work!", cherry)
		return
	}
	fmt.Println("Building an exe file from the Go code...")
	cmd := exec.Command("go", "build", "-o", "GoWeb.exe") // purposely using specific name
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Something happened when trying to build the codebase!
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		//fmt.Println("Something happened when trying to build the codebase! \n", err)
		return
	}
	fmt.Println("Running built exe file...")
	cmd = exec.Command(".\\GoWeb.exe") // running the exe to produce a local copy of the webpage
	output, err = cmd.CombinedOutput()
	if err != nil {
		// Something happened when trying to run the built command!
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println("GoWeb activated! Begin Testing...") // will be "connecting" using rod within the tests themselves
	m.Run()
	fmt.Println("Testing Complete!")

	//NOTE: GoWeb.exe isn't stopped (TODO), make sure you delete it before running more tests!
}

func TestPageRunning(t *testing.T) {

	browser := rod.New().MustConnect() // opens up the default browser
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	browser.MustPage("http://localhost:80/") // creates a page from browser, connects to localhost
}

func TestProperLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1
	// open up localhost as above
	// find input 1, 2 -> username pass
	// do input using user1, pass1
	// find login button
	// click
	// check page, make sure its on the non-login/non-error page (anything else is good)- error here

	// be sure to defer pageclose
}
func TestImproperLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1
	// open up localhost as above
	// find input 1, 2 -> username pass
	// do input using user0, pass0 (or funny names, just not the actual ones)
	// find login button
	// click
	// check page, make sure its on the login/error page (anything else is BAD)- error here

	// be sure to defer pageclose
}

func TestThetests(t *testing.T) {
	t.Errorf("Used to test TestMain's functions, make sure to comment out once done!")
}
