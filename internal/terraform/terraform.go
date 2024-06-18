package terraform

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// terraform cmd for plan and show output to json
func Tfcmd() interface{} {
	cmd := exec.Command("terraform", "plan", "-out=plan")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ERROR: Terraform plan completed with error. Please fix it")
		os.Exit(1)
	}
	cmd2 := exec.Command("terraform", "show", "-json", "plan")
	output, _ := cmd2.Output()
	var state interface{}
	json.Unmarshal(output, &state)
	defer os.Remove("./plan")
	return state
}
