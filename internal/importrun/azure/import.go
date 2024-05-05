package azure

import (
	"fmt"
	"os"
	"strings"
)

// function to call respective resource type to get their id
func Azure(subs_id string, res_type string, address string, config map[string]interface{}) int {
	file, err := os.OpenFile("./import.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	switch {
	// azurerm_resource_group
	case res_type == "azurerm_resource_group":
		if config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subs_id, config["name"].(string))
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

		// azurerm_virtual_network
	case res_type == "azurerm_virtual_network":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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

	// azurerm_subnet
	case res_type == "azurerm_subnet":
		if config["resource_group_name"] != nil && config["virtual_network_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", subs_id, config["resource_group_name"].(string), config["virtual_network_name"].(string), config["name"].(string))
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

		// azurerm_route_table
	case res_type == "azurerm_route_table":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/routeTables/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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

		// azurerm_network_security_group
	case res_type == "azurerm_network_security_group":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkSecurityGroups/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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

	// azurerm_network_security_rule
	case res_type == "azurerm_network_security_rule":
		if config["resource_group_name"] != nil && config["network_security_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkSecurityGroups/%s/securityRules/%s", subs_id, config["resource_group_name"].(string), config["network_security_group_name"].(string), config["name"].(string))
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

		// azurerm_virtual_network_peering
	case res_type == "azurerm_virtual_network_peering":
		if config["resource_group_name"] != nil && config["virtual_network_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/virtualNetworkPeerings/%s", subs_id, config["resource_group_name"].(string), config["virtual_network_name"].(string), config["name"].(string))
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

	// azurerm_subnet_route_table_association
	case res_type == "azurerm_subnet_route_table_association":
		if config["subnet_id"] != nil {
			id := fmt.Sprintf("%s", config["subnet_id"])
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

	// azurerm_subnet_network_security_group_association
	case res_type == "azurerm_subnet_network_security_group_association":
		if config["subnet_id"] != nil {
			id := fmt.Sprintf("%s", config["subnet_id"])
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

		// azurerm_management_group
	case res_type == "azurerm_management_group":
		if config["name"] != nil {
			id := fmt.Sprintf("/providers/Microsoft.Management/managementGroups/%s", config["name"].(string))
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

		// azurerm_storage_account
	case res_type == "azurerm_storage_account":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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

		// azurerm_virtual_machine or azurerm_linux_virtual_machine or azurerm_windows_virtual_machine
	case res_type == "azurerm_virtual_machine" || res_type == "azurerm_linux_virtual_machine" || res_type == "azurerm_windows_virtual_machine":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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

		// azurerm_virtual_desktop_host_pool
	case res_type == "azurerm_virtual_desktop_host_pool":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/hostPools/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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
		// azurerm_virtual_desktop_application_group
	case res_type == "azurerm_virtual_desktop_application_group":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/applicationGroups/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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
		// azurerm_virtual_desktop_workspace
	case res_type == "azurerm_virtual_desktop_workspace":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/workspaces/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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
		// azurerm_virtual_desktop_workspace_application_group_association
	case res_type == "azurerm_virtual_desktop_workspace_application_group_association":
		if config["workspace_id"] != nil && config["application_group_id"] != nil {
			id := fmt.Sprintf("%s|%s", config["workspace_id"].(string), config["application_group_id"].(string))
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
		// azurerm_virtual_desktop_scaling_plan
	case res_type == "azurerm_virtual_desktop_scaling_plan":
		if config["resource_group_name"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/scalingPlans/%s", subs_id, config["resource_group_name"].(string), config["name"].(string))
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
		// azurerm_virtual_desktop_application
	case res_type == "azurerm_virtual_desktop_application":
		if config["resource_group_name"] != nil && config["application_group_id"] != nil && config["name"] != nil {
			id := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/applicationGroups/%s/applications/%s", subs_id, config["resource_group_name"].(string), config["application_group_id"].(string), config["name"].(string))
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
	default:
		return 0
	}

}
