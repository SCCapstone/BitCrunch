package testing

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/go-rod/rod"
)

var TestSignupCheck = false // for making a correct sign-in
// https://github.com/go-rod/rod/blob/master/examples_test.go

func TestMain(m *testing.M) {
	fmt.Println("Please note: in order to properly test, please consult the Readme! Otherwise, the browser will NOT connect properly!")
	fmt.Println("Building and Activating GoWeb...")
	fmt.Println("Making sure that we are in the correct directory...")
	cherry := os.Chdir("../../")
	if cherry != nil {
		// Something happened when trying to change dir!
		fmt.Println("Command chdir didn't execute!", cherry)
		return
	}
	fmt.Println("Building an exe file from the Go code...")
	cmd := exec.Command("go", "build", "-o", "GoWeb.exe") // purposely using specific name
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Something happened when trying to build the codebase!
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println("Running built exe file...")
	//func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
	cmd = exec.Command(".\\GoWeb.exe") // running the exe to produce a local copy of the webpage
	output, err = cmd.CombinedOutput()
	if err != nil {
		// something happenned on executing the exe!
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println("GoWeb activated! Begin Testing...") // will be "connecting" using rod within the tests themselves
	m.Run()                                          // All ""selected"" tests are run here.
	fmt.Println("Testing Complete!")
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		fmt.Println("GoWeb cannot be killed. The uprising has begun. (Try killing it manually using Task Manager/System Monitor)")
		return
	}
	fmt.Println("GoWeb has been closed.") // undecided as what to place here
	// TODO run a check for any db files created, txt files in devices delete them
	fmt.Println("Cleaning up created files...")
	err = os.Remove("floors.db")
	if err != nil {
		fmt.Println("There's an error in finding and deleting this file!", err)
		// no return call
	}
	err = os.Remove("users.db")
	if err != nil {
		fmt.Println("There's an error in finding and deleting this file!", err)
		// no return call
	}
	err = RemoveContents("devices")
	if err != nil {
		fmt.Println("An error was raised when cleaning the devices folder, ", err)
	}

	fmt.Println("All Behavorial Tests Completed!")
}
func TestSignUp(t *testing.T) {
	/// on signin page
	browser := rod.New().MustConnect()
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	page := browser.MustPage("http://localhost:5000/") // creates a page from browser, connects to localhost
	// click signup
	page.MustElement("form#/u/register").MustElement("input.signup.moveup").MustClick() // check form#/u/ ... for correctness
	// fill out all inputs incorrectly
	elUser := page.MustElement("form#/u/register").MustElement("input#username") // FIX: its <__ action > that is being looked for

	elUser.MustInput("user1")

	elPass := page.MustElement("form#/u/register").MustElement("input#password")

	elPass.MustInput("Passsword1!")
	elPassConfirm := page.MustElement("form#/u/register").MustElement("input#confirm_password") // FIX: its <__ action > that is being looked for

	elPassConfirm.MustInput("PassPassword1!")

	elEmail := page.MustElement("form#/u/register").MustElement("input#email")

	elEmail.MustInput("realemail@gmail.com")
	// click sign up
	page.MustElement("form").MustElement("div").MustClick() // this clicks on the first form/div element. it happens to be the submit bttn
	// check URL for failure
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http://localhost:5000/map$/")
	// I could also check the db file that's populated for anything that pops up
	if page == nil { // couldn't find the right url
		t.Errorf("The signup wasn't successfull!")
		return
	}
	TestSignupCheck = true //signal globally that Signup has been completed
}
func TestImproperSignup(t *testing.T) {
	// on signin page
	browser := rod.New().MustConnect()
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	page := browser.MustPage("http://localhost:5000/") // creates a page from browser, connects to localhost
	// click signup
	page.MustElement("form#/u/register").MustElement("input.signup.moveup").MustClick() // check form#/u/ ... for correctness
	// fill out all inputs incorrectly
	elUser := page.MustElement("form#/u/register").MustElement("input#username") // FIX: its <__ action > that is being looked for

	elUser.MustInput("literally anything")

	elPass := page.MustElement("form#/u/register").MustElement("input#password")

	elPass.MustInput("dumbdumbwithnocapitalsnumbersorfunnysymbols")
	elPassConfirm := page.MustElement("form#/u/register").MustElement("input#confirm_password") // FIX: its <__ action > that is being looked for

	elPassConfirm.MustInput("somthingelse")

	elEmail := page.MustElement("form#/u/register").MustElement("input#email")

	elEmail.MustInput("FAKEemail(with a space and symbols(bad email things)) @ gfail.cym")
	// click sign up
	page.MustElement("form").MustElement("div").MustClick() // this clicks on the first form/div element. it happens to be the submit bttn
	// check URL for failure
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http://localhost:5000/map$/")
	// no regex. only fixed fit >:(
	if page != nil { // couldn't find the right url
		t.Errorf("The signup was successfull! It really shouldn't be!")
		return
	}

}

func TestPageRunning(t *testing.T) { // used for making sure the webapp loads correctly
	browser := rod.New().MustConnect() // opens up the default browser
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}() // technically a lambda function btw
	browser.MustPage("http://localhost:5000/") // creates a page from browser, connects to localhost
}

func TestProperLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1
	// open up localhost as above
	if !TestSignupCheck {
		TestSignUp(t) // used for making sure there's good sign in credentials
	}
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
	elUser := page.MustElement("form").MustElement("input#username") // FIX: its <__ action > that is being looked for

	elUser.MustInput("user1")

	elPass := page.MustElement("form").MustElement("input#password")

	elPass.MustInput("PassPassword1!")
	// find login button, click it
	page.MustElement("form").MustElement("div").MustClick() // this clicks on the first form/div element. it happens to be the submit bttn
	// check page, make sure its on the non-login/non-error page (anything else is good)
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http://localhost:5000/map$/")
	// got rid of the regex. hardcoded URL now
	if page == nil { // couldn't find the right url
		t.Errorf("The login was not successfull! There is an issue on typing login!")
		return
	}
}
func TestImproperLogin(t *testing.T) { // opens up the domain and attempts to login using oxymoron and Ferroseed4732#!
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
	elUser := page.MustElement("form").MustElement("input#username") // FIX: its <__ action > that is being looked for

	elUser.MustInput("oxymoron")

	elPass := page.MustElement("form").MustElement("input#password")

	elPass.MustInput("Ferroseed4732#!")
	// find login button
	// click it
	page.MustElement("form").MustElement("div").MustClick()
	// check page, make sure its on the non-login/non-error page (anything else is good)- error here
	currpages, err := browser.Pages()
	if err != nil { // check to see if the page was rendered at all
		t.Errorf("There was an issue rendering the map/floors page!")
		return
	}
	page, err = currpages.FindByURL("/^http://localhost:5000/u/login$/")
	// god i hate regex so much. no more regex
	if page != nil { // couldn't find the right url
		t.Errorf("The login was successfull! However, it should'nt be, as these credentials are not correct")
		return
	}
}

func TestLogout(t testing.T) {
	// navigate to main page (see TestLogin)
	// click Logout
	// check domain with FindByUrl
	// open up localhost as above
	browser := rod.New().MustConnect()
	defer func() {
		_, err := browser.Pages()
		if err != nil { // check to see if the page was rendered at all
			t.Errorf("There was an issue rendering the webapp!")
		}
		browser.MustClose() // On panic (and end), close the browser
	}()
	page := browser.MustPage("http://localhost:5000/")               // creates a page from browser, connects to localhost
	elUser := page.MustElement("form").MustElement("input#username") // FIX: its <__ action > that is being looked for
	elUser.MustInput("user1")
	elPass := page.MustElement("form").MustElement("input#password")
	elPass.MustInput("pass1")
	// find login button, click it
	page.MustElement("form").MustElement("div").MustClick()
	// now logged in, in the main welcome page

	page.MustElement("div.parent").MustElement("div.grid-item.item3").MustElement("button.style1_button").MustClick()
	// logout has been clicked, now on logout_modal
	page.MustElement("div#logout_modal").MustElement("form").MustElement("input.danger_button").MustClick() // click yes
	currpages, _ := browser.Pages()
	page, _ = currpages.FindByURL("/^http://localhost:5000/$/")
	// no regex. only fixed fit >:( regex is dumb >:(
	if page != nil { // couldn't find the right url
		t.Errorf("The signup was successfull! It really shouldn't be!")
		return
	}

}

// helper function used to clean a directory
// https://stackoverflow.com/questions/33450980/how-to-remove-all-contents-of-a-directory-using-golang
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func TestThetests(t *testing.T) {
	t.Errorf("Used to test TestMain's functions, make sure to comment out once done!")
}
*/

// login
// /map
// settings, logout, add layer, possible layers
// / view_layer
// everything in /map, + edit delete add device buttons  + add device
// / view_device
// edit save cancel delete (run scripts?)
