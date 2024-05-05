package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tfimport/internal/contains"
	"github.com/tfimport/internal/importcheck"
	"github.com/tfimport/internal/importrun/azure"
	"github.com/tfimport/internal/importrun/google"
	"github.com/tfimport/internal/provider"
	"github.com/tfimport/internal/terraform"
)

// Main function to start the tool
func main() {

	if len(os.Args[:]) <= 1 {
		panic("Pass arguement setup or plan")
	}
	state := terraform.Tfcmd()
	if state.(map[string]interface{})["errored"].(bool) {
		panic("ERROR: Terraform plan completed with error. Please fix it")
	}
	if os.Args[1] == "setup" {
		provider.Provider_config(state)
	} else if os.Args[1] == "plan" {
		pjs, _ := os.ReadFile("./provider.json")
		var provider_json_file map[string]string
		op2 := json.Unmarshal(pjs, &provider_json_file)
		if op2 != nil {
			log.Fatal(op2)
		}
		var import_drive string = "proceed"
		for import_drive == "proceed" {

			var each_base_import_count int = 0
			var total_base_import_count int = 0
			state := terraform.Tfcmd()
			if state.(map[string]interface{})["resource_changes"] != nil {
				resource_list := state.(map[string]interface{})["resource_changes"].([]interface{})
				for res_list := 0; res_list < len(resource_list); res_list++ {
					res_type := resource_list[res_list].(map[string]interface{})["type"].(string)
					address := resource_list[res_list].(map[string]interface{})["address"].(string)
					mode := resource_list[res_list].(map[string]interface{})["mode"].(string)
					name := resource_list[res_list].(map[string]interface{})["name"].(string)
					provider_name := resource_list[res_list].(map[string]interface{})["provider_name"].(string)
					change := resource_list[res_list].(map[string]interface{})["change"].(map[string]interface{})["actions"].([]interface{})
					after := resource_list[res_list].(map[string]interface{})["change"].(map[string]interface{})["after"].(map[string]interface{})
					if contains.Contains(change, "create") && mode == "managed" {
						if provider_name == "registry.terraform.io/hashicorp/azurerm" {
							pck := provider.Provider_config_key(address, res_type, mode, name, state)
							subs := provider_json_file[pck]
							each_base_import_count = azure.Azure(subs, res_type, address, after)
							total_base_import_count = total_base_import_count + each_base_import_count
						}
						if provider_name == "registry.terraform.io/hashicorp/google" {
							each_base_import_count = google.Google(res_type, address, after)
							total_base_import_count = total_base_import_count + each_base_import_count
						}
					}
				}
			}
			if total_base_import_count != 0 {
				_ = importcheck.Import_check()
				import_drive = "proceed"
			} else {
				result := importcheck.Import_check()
				if result == "no_err" {
					import_drive = "done"
				} else {
					import_drive = "proceed"
				}
			}
		}
		importcheck.Remove_commented_lines()
	} else {
		panic("Pass arguement setup or plan")
	}
}
