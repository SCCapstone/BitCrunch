package testing

import (
	"fmt"
	"testing"

	models "github.com/SCCapstone/BitCrunch/models"

	"github.com/go-rod/rod"
)

func TestMain(m *testing.M) {
	fmt.Println("placeholder")
	browser := rod.New().MustConnect()
	defer browser.MustClose() // makes sure the browser closes once Tests are complete
}

func TestBrowser(t *testing.T) {
	newfloorList := models.GetAllFloors()
	fmt.Println(newfloorList[0]) // placeholder

}
