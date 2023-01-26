package testing

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

// https://github.com/go-rod/rod/blob/master/examples_test.go

func TestMain(m *testing.M) {
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
	/*
		browser := rod.New().MustConnect() // opens up the default browser
		fmt.Printf("browser")
		defer browser.MustClose()              // makes sure the browser closes once Tests are complete
		page1 := browser.MustPage("localhost") // creates a page from browser
		fmt.Printf("page1")
		fmt.Println(page1)
		fmt.Printf("defer close")
	*/
	openbrowser("localhost")
	fmt.Println("GoWeb activated! Begin Testing...") // will be "connecting" using rod within the tests themselves
	m.Run()
	fmt.Println("Testing Complete!")

	//NOTE: GoWeb.exe isn't removed by the code (TODO), make sure you delete it before running tests!
}

func TestPageRunning(t *testing.T) {

}

func TestLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1

}

// Helper function to open correct browser for testee's machine
// thanks to https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
