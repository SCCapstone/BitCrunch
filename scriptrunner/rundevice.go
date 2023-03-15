package scriptrunner

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// function takes in the exact name of the script, searches for it in assets
// After found, runs each line through the OS's commandline
// Importantly, any line beginning with # are discarded; they are considered comments
func RunFromScript(filename string, targetIP string) (error, string) {
	IPaddress := "<IPADDRESS>" // for better usage across many files and devices
	outputAll := ""
	// static\assets\pingscript.txt
	betterFilename := "static\\assets\\" + filename // gotta get the right file!
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
		IPfound, index := contains(line, IPaddress)
		if IPfound && index != -1 {
			line[index] = targetIP // place the parameter IPaddress properly into the script
		}
		// make sure to test this!
		if len(line) == 0 {
			continue
		}
		if string(line[0]) == ("#") || line == nil {
			continue // This line is a comment or empty! don't do anything!
		}
		var cmd *exec.Cmd
		//https://stackoverflow.com/questions/22781788/how-could-i-pass-a-dynamic-set-of-arguments-to-gos-command-exec-command
		//- for arguments in exec.Command
		if len(line) == 1 {
			// just solo exec.Command(line)

			cmd = exec.Command(line[0])
			//fmt.Println("solo", cmd)
		} else {
			// split up into exec.Command(line[0], line[1:len(line)-1])
			var allArgs string
			for index, argument := range line {
				if index == 0 {
					continue
				}
				allArgs = allArgs + argument
			}
			cmd = exec.Command(line[0], allArgs)
			//fmt.Println("duo", cmd)
		}
		if cmd == nil {
			continue
		}
		output, err := cmd.CombinedOutput() // this works, bc command is NEVER empty! (kinda)
		// just in case, there's a check right above
		if err != nil {
			file.Close()
			return err, string(output)
		}

		outputAll += string(output) // move the output into the full output
	}
	file.Close()
	return nil, outputAll
}

func contains(fullList []string, soloString string) (bool, int) {
	for index, a := range fullList {
		if a == soloString {
			return true, index
		}
	}
	return false, -1
}
