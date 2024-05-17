package google

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// function to call respective resource type to get their id
func Google(data map[string]interface{}, res_type string, address string, config map[string]interface{}) int {
	file, err := os.OpenFile("./import.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	if value, ok := data[res_type]; ok {
		pattern := regexp.MustCompile("{{(.*?)}}")
		matches := pattern.FindAllStringSubmatch(value.(map[string]interface{})["id"].(string), -1)
		var params []string
		for _, match := range matches {
			if len(match) > 1 {
				params = append(params, match[1])
			}
		}
		count := len(params)
		for _, iv := range params {
			if _, ok = config[iv]; ok {
				count--
			}
		}
		if count == 0 {
			id := value.(map[string]interface{})["id"].(string)
			for _, iv := range params {
				id = strings.ReplaceAll(id, "{{"+iv+"}}", config[iv].(string))
			}
			for ik, iv := range value.(map[string]interface{})["change"].(map[string]interface{}) {
				id = strings.ReplaceAll(id, ik, iv.(string))
			}
			address_check, _ := os.ReadFile("./import.tf")
			if !strings.Contains(string(address_check), address) {
				Import := fmt.Sprintf(`
import {
  id = "%s"
  to = %s
}
	
`, id, address)
				_, err = file.WriteString(Import)
				if err != nil {
					fmt.Println("Error writing to import.tf file:", err)
					return 0
				}
				return 1
			}
			return 0

		}
		return 0
	}
	return 0

}
