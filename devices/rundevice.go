package devices

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// function takes in the exact name of the script, searches for it in assets
// After found, runs each line through the OS's commandline
// Importantly, any line beginning with # are discarded; they are considered comments
func RunFromScript(filename string) (error, string) {
	// gotta go to the right file!
	outputAll := ""
	betterFilename := "/static/assets/" + filename
	// the paaram should be the file name, no txt
	file, err := os.Open(betterFilename)
	if err != nil {
		return err, outputAll
		//fmt.Print("do something")
	}
	fileScanner := bufio.NewScanner((file))
	fileScanner.Split(bufio.ScanLines) // this splits the file by newline (super easy)
	for fileScanner.Scan() {           // this moves until end of file
		//filescanner.Text() = current line
		line := strings.Fields(fileScanner.Text()) // split by space, tab, newline, etc.
		if string(line[0]) == ("#") || line == nil {
			continue // This line is a comment or empty! don't do anything!
		}
		var cmd *exec.Cmd
		if len(line) == 1 {
			// just solo exec.Command(line)
			cmd = exec.Command(line[0])
		} else {
			// split up into exec.Command(line[0], line[1:len(line)-1])
			cmd = exec.Command(line[0], line[1:len(line)-1]...) // literally magic ...
		}
		if cmd == nil {
			continue
		}
		output, err := cmd.CombinedOutput() // this works, even w empty cmd!
		//https://stackoverflow.com/questions/22781788/how-could-i-pass-a-dynamic-set-of-arguments-to-gos-command-exec-command
		// :) - for arguments in exec.Command

		if err != nil {
			file.Close()
			return err, string(output)
		}
	}

	if err != nil {
		file.Close()
		fmt.Print("do something")
	}
	return nil, outputAll
}
