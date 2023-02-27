package testing

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Running All Tests...")
	m.Run()
	fmt.Println("All Unit Tests Completed!")
}
