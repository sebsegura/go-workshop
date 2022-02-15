package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()

	// rc 0 means we've passed
	// and CoverMode will be non-empty if we run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0 {
			fmt.Println("Test passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}
