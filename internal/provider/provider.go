package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Provider_config(state interface{}) {
	provider_map := make(map[string]string)
	if configuration, ok := state.(map[string]interface{})["configuration"]; ok {
		if provider_config, ok := configuration.(map[string]interface{})["provider_config"]; ok {
			for key, value := range provider_config.(map[string]interface{}) {
				if value.(map[string]interface{})["full_name"] == "registry.terraform.io/hashicorp/azurerm" {
					provider_map[key] = "Update Subs id"
				}
			}
			jsonData, err_jd := json.MarshalIndent(provider_map, "", "	")
			if err_jd != nil {
				fmt.Println("Error marshaling JSON:", err_jd)
				return
			}
			err_wf := os.WriteFile("provider.json", jsonData, 0644)
			if err_wf != nil {
				fmt.Println("Error writing JSON to file:", err_wf)
				return
			}
		}
	}
}

// provider config key extraction
func Provider_config_key(address string, resource_type string, mode string, name string, state interface{}) string {
	count := strings.Count(address, ".")
	var pck string
	list_pck := state.(map[string]interface{})["configuration"].(map[string]interface{})["root_module"]
	for dep := 1; dep < count; dep += 2 {
		list_pck = list_pck.(map[string]interface{})["module_calls"].(map[string]interface{})[strings.Split(strings.Split(address, ".")[dep], "[")[0]].(map[string]interface{})["module"]
	}
	list_pck_converted := list_pck.(map[string]interface{})["resources"].([]interface{})
	for i := 0; i < len(list_pck_converted); i++ {
		if list_pck_converted[i].(map[string]interface{})["mode"] == mode && list_pck_converted[i].(map[string]interface{})["type"] == resource_type && list_pck_converted[i].(map[string]interface{})["name"] == name {
			pck = list_pck_converted[i].(map[string]interface{})["provider_config_key"].(string)
			break
		}
	}
	return pck
}
