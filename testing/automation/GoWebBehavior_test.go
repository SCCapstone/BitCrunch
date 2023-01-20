package testing

import (
	"fmt"
	"os/exec"
	"testing"

	models "github.com/SCCapstone/BitCrunch/models"

	"github.com/go-rod/rod"
)

// https://github.com/go-rod/rod/blob/master/examples_test.go

func TestMain(m *testing.M) {
	fmt.Println("Building and Activating GoWeb...")
	fmt.Println("Building an exe file from the Go code...")
	cmd := exec.Command("go", "build -o GoWeb.exe") // purposely using specific name
	err := cmd.Run()
	if err != nil {
		// Something happened when trying to build the codebase!
		fmt.Println("Something happened when trying to build the codebase! \n", err)
		return
	}
	fmt.Println("Running built exe file...")
	cmd = exec.Command(".\\GoWeb.exe") // running the exe to produce a local copy of the webpage
	err = cmd.Run()
	if err != nil {
		// Something happened when trying to run the built command!
		fmt.Println("Something happened when trying to run the built command! \n", err)
		return
	}
	browser := rod.New().MustConnect() // opens up the default browser
	defer browser.MustClose()          // makes sure the browser closes once Tests are complete

	fmt.Println("GoWeb activated! Begin Testing...")
	m.Run()
	fmt.Println("Testing Complete!")
}

func TestBrowser(t *testing.T) {
	newfloorList := models.GetAllFloors()
	fmt.Println(newfloorList[0]) // placeholder

}

func TestLogin(t *testing.T) { // opens up the domain and attempts to login using user1 and pass1

}
