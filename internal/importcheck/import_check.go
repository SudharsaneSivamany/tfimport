package importcheck

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Import_check() string {
	cmd := exec.Command("terraform", "plan", "-no-color")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// Input paragraph
		errlines := strings.Split(string(out), "\n")

		// Regular expression pattern for start and end markers
		startPattern := `While attempting to import an existing object to`
		endPattern := `", the provider detected that no object exists with the given id. Only`

		var errorLinesnumber []int
		var concatErrorLines []string
		var addressError []string
		var unknown_error_count int = 0

		for i := 0; i < len(errlines); i++ {
			if strings.Contains(errlines[i], startPattern) {
				errorLinesnumber = append(errorLinesnumber, i)
			}
		}
		errorLinesnumber = append(errorLinesnumber, len(errlines)-1)
		for i := 0; i < len(errorLinesnumber)-1; i++ {
			var temp string = ""
			for j := errorLinesnumber[i]; j < errorLinesnumber[i+1]; j++ {
				temp = temp + errlines[j] + " "
			}
			concatErrorLines = append(concatErrorLines, temp)
		}
		for i := 0; i < len(concatErrorLines); i++ {
			line := strings.Split(concatErrorLines[i], startPattern)[1]
			line = strings.Split(line, endPattern)[0]
			line = strings.ReplaceAll(line, `[\"`, `[\`)
			line = strings.ReplaceAll(line, `\"]`, `\]`)
			line = strings.Split(line, `"`)[1]
			line = strings.ReplaceAll(line, `[\`, `["`)
			line = strings.ReplaceAll(line, `\]`, `"]`)
			addressError = append(addressError, line)
			unknown_error_count += 1
		}
		file, _ := os.ReadFile("./import.tf")
		lines := strings.Split(string(file), "\n")
		lines_count := len(lines)
		new_file, err1 := os.OpenFile("./import.tf", os.O_TRUNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			fmt.Println("Error opening file:", err1)
			return "err"
		}
		defer new_file.Close()
		for _, value := range addressError {
			for i := 0; i < lines_count; i++ {
				if strings.Contains(string(lines[i]), value) {
					// Comment out the line
					lines[i-2] = "#commented# " + string(lines[i-2])
					lines[i-1] = "#commented# " + string(lines[i-1])
					lines[i] = "#commented# " + string(lines[i])
					lines[i+1] = "#commented# " + string(lines[i+1])
				}
			}
		}
		for i := 0; i < lines_count; i++ {
			_, err3 := new_file.WriteString(lines[i])
			if err3 != nil {
				fmt.Println("Error writing to import.tf file:", err3)
				return "err"
			}
			new_file.WriteString("\n")
		}
		if unknown_error_count == 0 {
			panic("ERROR UNKNOWN: Please run terraform plan")
		} else {
			return "err"
		}

	} else {
		return "no_err"
	}
}

func Remove_commented_lines() {
	file, _ := os.ReadFile("./import.tf")
	lines := strings.Split(string(file), "\n")
	lines_count := len(lines)
	new_file, err1 := os.OpenFile("./import.tf", os.O_TRUNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("Error opening file:", err1)
		return
	}
	defer new_file.Close()
	for i := 0; i < lines_count; i++ {
		if !strings.Contains(string(lines[i]), "#commented#") {
			_, err3 := new_file.WriteString(lines[i])
			if err3 != nil {
				fmt.Println("Error writing to import.tf file:", err3)
				return
			}
			new_file.WriteString("\n")
		}
	}

}
