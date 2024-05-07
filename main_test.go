package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	want := "welco"
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	output, _ := io.ReadAll(r)
	os.Stdout = old
	fmt.Println(string(output))
	if want != strings.TrimSpace(string(output)) {
		fmt.Println("Unexpected output. Got:", string(output), "Want:", want)
		os.Exit(1) // Exit with non-zero status to indicate test failure
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

// old := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	// Call the Trial function
// 	Trial(inputFile, inputString)

// 	// Read printed output
// 	w.Close()
// 	captured, _ := ioutil.ReadAll(r)
// 	os.Stdout = old
