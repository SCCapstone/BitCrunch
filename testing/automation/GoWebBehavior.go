package testing

import (
	"fmt"
	"testing"

	main "github.com/SCCapstone/BitCrunch"

	"github.com/go-rod/rod"
)

func TestMain(m *testing.M) {
	fmt.Println("placeholder")
	browser := rod.New().MustConnect()
	defer browser.MustClose() // makes sure the browser closes once Tests are complete
}

func TestBrowser(t *testing.T) {
	floorList := main.GetAllFloors()

}
