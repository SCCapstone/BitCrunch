package main

import (
	"fmt"
	"testing"

	"github.com/go-rod/rod"
)

func TestMain(m *testing.M) {
	fmt.Println("placeholder")
	browser := rod.New().MustConnect()
	defer browser.MustClose()
}
