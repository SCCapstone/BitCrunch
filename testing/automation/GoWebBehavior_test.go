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
	//func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
	cmd = exec.Command(".\\GoWeb.exe") // running the exe to produce a local copy of the webpage
	output, err = cmd.CombinedOutput()
	if err != nil {
		// Something happened when trying to run the built command!
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println("GoWeb activated! Begin Testing...") // will be "connecting" using rod within the tests themselves
	m.Run()                                          // All ""selected"" tests are run here.
	fmt.Println("Testing Complete!")
	err = cmd.Process.Kill()
	if err != nil {
		// Something happened when trying to kill GoWeb.
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		fmt.Println("GoWeb cannot be killed. The uprising has begun. (Try killing it manually using Task Manager/System Monitor)")
		return
	}
	fmt.Println("GoWeb has been closed.") // undecided as what to place here
}

func TestPageRunning(t *testing.T) {

	browser := rod.New().MustConnect() // opens up the default browser
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}() // technically a lambda function btw
	browser.MustPage("http://localhost:80/") // creates a page from browser, connects to localhost
}

func TestProperLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1
	// open up localhost as above
	browser := rod.New().MustConnect()
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	page := browser.MustPage("http://localhost:5000/") // creates a page from browser, connects to localhost

	// https://go-rod.github.io/#/input

	// find input 1, 2 -> username pass
	// do input using user1, pass1
	//fmt.Println("going for username look")
	elUser := page.MustElement("form").MustElement("input#username") // FIX: its <__ action > that is being looked for
	//fmt.Println("Found", elUser, "input element for \"username\". ")

	//fmt.Println("going for username input")
	elUser.MustInput("user1")

	//fmt.Println(elUser.MustText()) // use MustText to get the text
	//fmt.Println("going for pass look")

	elPass := page.MustElement("form").MustElement("input#password")
	//fmt.Println("going for pass input")

	elPass.MustInput("pass1")
	//fmt.Println(elPass.MustText()) // use MustText to get the text
	// find login button
	// click it
	page.MustElement("form").MustElement("div").MustClick()
	// check page, make sure its on the non-login/non-error page (anything else is good)- error here
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http:\\/\\/\\w+(\\.\\w+)*(:[0-9]+)?\\/?(\\/[.\\w]*)*$/")
	// god i hate regex so much. this checks for http(anything)/(anything) rn. find proper directory soon!
	if page == nil { // couldn't find the right url
		t.Errorf("The login was not successfull! There is an issue on typing login!")
		return
	}
	// be sure to defer pageclose
}
func TestImproperLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1
	// open up localhost as above
	browser := rod.New().MustConnect()
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	page := browser.MustPage("http://localhost:5000/") // creates a page from browser, connects to localhost

	// https://go-rod.github.io/#/input

	// find input 1, 2 -> username pass
	// do input using user1, pass1
	//fmt.Println("going for username look")
	elUser := page.MustElement("form").MustElement("input#username") // FIX: its <__ action > that is being looked for
	//fmt.Println("Found", elUser, "input element for \"username\". ")

	//fmt.Println("going for username input")
	elUser.MustInput("oxymoron")

	//fmt.Println(elUser.MustText()) // use MustText to get the text
	//fmt.Println("going for pass look")

	elPass := page.MustElement("form").MustElement("input#password")
	//fmt.Println("going for pass input")

	elPass.MustInput("ferroseed4732#!")
	//fmt.Println(elPass.MustText()) // use MustText to get the text
	// find login button
	// click it
	page.MustElement("form").MustElement("div").MustClick()
	// check page, make sure its on the non-login/non-error page (anything else is good)- error here
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http:\\/\\/\\w+(\\.\\w+)*(:[0-9]+)?\\/?(\\/[.\\w]*)*$/")
	// god i hate regex so much
	if page != nil { // couldn't find the right url
		t.Errorf("The login was successfull! However, it should'nt be, as these credentials are not correct")
		return
	}
	// be sure to defer pageclose
}

/*
func TestThetests(t *testing.T) {
	t.Errorf("Used to test TestMain's functions, make sure to comment out once done!")
}
*/
