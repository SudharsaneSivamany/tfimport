package importcheck

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Import_check() string {
	cmd := exec.Command("terraform", "plan")
	out, err := cmd.CombinedOutput()
	if err != nil {
		errors := strings.Split(string(out), "\n")
		startPattern := `While attempting to import an existing object to`
		endPattern := `the provider detected that no object exists with the given id. Only`
		foundStart := false
		foundEnd := false
		var addressError []string
		var unknown_error_count int = 0
		for _, line := range errors {
			// Check if the line matches the startPattern
			if !foundStart && regexp.MustCompile(startPattern).MatchString(line) {
				foundStart = true
				continue
			}

			// Check if the line matches the endPattern
			if foundStart && !foundEnd && regexp.MustCompile(endPattern).MatchString(line) {
				foundEnd = true
				break
			}

			// Output the line if it's between startPattern and endPattern
			if foundStart && !foundEnd {
				line = strings.Split(line, `",`)[0]
				line = strings.ReplaceAll(line, `[\"`, `[\`)
				line = strings.ReplaceAll(line, `\"]`, `\]`)
				line = strings.Split(line, `"`)[1]
				line = strings.ReplaceAll(line, `[\`, `["`)
				line = strings.ReplaceAll(line, `\]`, `"]`)
				addressError = append(addressError, line)
				foundStart = false
				foundEnd = false
				unknown_error_count += 1
			}
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
