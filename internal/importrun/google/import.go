package google

import (
	"fmt"
	"os"
	"strings"
)

// function to call respective resource type to get their id
func Google(res_type string, address string, config map[string]interface{}) int {
	file, err := os.OpenFile("./import.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	switch {
	// google_compute_network
	case res_type == "google_compute_network":
		if config["project"] != nil && config["name"] != nil {
			id := fmt.Sprintf("projects/%s/global/networks/%s", config["project"].(string), config["name"].(string))
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
	// google_compute_subnetwork
	case res_type == "google_compute_subnetwork":
		if config["project"] != nil && config["region"] != nil && config["name"] != nil {
			id := fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", config["project"].(string), config["region"].(string), config["name"].(string))
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
		// google_api_gateway_api
	case res_type == "google_api_gateway_api":
		if config["project"] != nil && config["api_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s", config["project"].(string), config["api_id"].(string))
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
		// google_api_gateway_api_config
	case res_type == "google_api_gateway_api_config":
		if config["project"] != nil && config["api"] != nil && config["api_config_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", config["project"].(string), config["api"].(string), config["api_config_id"].(string))
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
		// google_api_gateway_api_config_iam_member
	case res_type == "google_api_gateway_api_config_iam_member":
		if config["project"] != nil && config["api"] != nil && config["api_config"] != nil && config["role"] != nil && config["member"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s %s %s", config["project"].(string), config["api"].(string), config["api_config"].(string), config["role"].(string), config["member"].(string))
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
		// google_api_gateway_api_config_iam_binding
	case res_type == "google_api_gateway_api_config_iam_binding":
		if config["project"] != nil && config["api"] != nil && config["api_config"] != nil && config["role"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s %s", config["project"].(string), config["api"].(string), config["api_config"].(string), config["role"].(string))
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
		// google_api_gateway_api_config_iam_policy
	case res_type == "google_api_gateway_api_config_iam_policy":
		if config["project"] != nil && config["api"] != nil && config["api_config"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", config["project"].(string), config["api"].(string), config["api_config"].(string))
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
		// google_api_gateway_api_iam_member
	case res_type == "google_api_gateway_api_iam_member":
		if config["project"] != nil && config["api"] != nil && config["role"] != nil && config["member"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s %s %s", config["project"].(string), config["api"].(string), config["role"].(string), config["member"].(string))
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
		// google_api_gateway_api_iam_binding
	case res_type == "google_api_gateway_api_iam_binding":
		if config["project"] != nil && config["api"] != nil && config["role"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s %s", config["project"].(string), config["api"].(string), config["role"].(string))
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
		// google_api_gateway_api_iam_policy
	case res_type == "google_api_gateway_api_iam_policy":
		if config["project"] != nil && config["api"] != nil {
			id := fmt.Sprintf("projects/%s/locations/global/apis/%s", config["project"].(string), config["api"].(string))
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
		// google_api_gateway_gateway
	case res_type == "google_api_gateway_gateway":
		if config["project"] != nil && config["region"] != nil && config["gateway_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", config["project"].(string), config["region"].(string), config["gateway_id"].(string))
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
		// google_api_gateway_gateway_iam_member
	case res_type == "google_api_gateway_gateway_iam_member":
		if config["project"] != nil && config["region"] != nil && config["gateway"] != nil && config["role"] != nil && config["member"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/gateways/%s %s %s", config["project"].(string), config["region"].(string), config["gateway"].(string), config["role"].(string), config["member"].(string))
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
		// google_api_gateway_gateway_iam_binding
	case res_type == "google_api_gateway_gateway_iam_binding":
		if config["project"] != nil && config["region"] != nil && config["gateway"] != nil && config["role"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/gateways/%s %s", config["project"].(string), config["region"].(string), config["gateway"].(string), config["role"].(string))
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
		// google_api_gateway_gateway_iam_policy
	case res_type == "google_api_gateway_gateway_iam_policy":
		if config["project"] != nil && config["region"] != nil && config["gateway"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", config["project"].(string), config["region"].(string), config["gateway"].(string))
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
		// google_alloydb_backup
	case res_type == "google_alloydb_backup":
		if config["project"] != nil && config["location"] != nil && config["backup_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/backups/%s", config["project"].(string), config["location"].(string), config["backup_id"].(string))
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
		// google_alloydb_cluster
	case res_type == "google_alloydb_cluster":
		if config["project"] != nil && config["location"] != nil && config["cluster_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", config["project"].(string), config["location"].(string), config["cluster_id"].(string))
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
		// google_alloydb_instance
	case res_type == "google_alloydb_instance":
		if config["project"] != nil && config["location"] != nil && config["cluster"] != nil && config["instance_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/instances/%s", config["project"].(string), config["location"].(string), config["cluster"].(string), config["instance_id"].(string))
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
		// google_alloydb_user
	case res_type == "google_alloydb_user":
		if config["project"] != nil && config["location"] != nil && config["cluster"] != nil && config["user_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/users/%s", config["project"].(string), config["location"].(string), config["cluster"].(string), config["user_id"].(string))
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
		// google_app_engine_application or google_app_engine_application_url_dispatch_rules
	case res_type == "google_app_engine_application" || res_type == "google_app_engine_application_url_dispatch_rules":
		if config["project"] != nil {
			id := fmt.Sprintf("%s", config["project"])
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
		// google_app_engine_domain_mapping
	case res_type == "google_app_engine_domain_mapping":
		if config["project"] != nil && config["domain_name"] != nil {
			id := fmt.Sprintf("apps/%s/domainMappings/%s", config["project"].(string), config["domain_name"].(string))
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
		// google_app_engine_firewall_rule
	case res_type == "google_app_engine_firewall_rule":
		if config["project"] != nil && config["priority"] != nil {
			id := fmt.Sprintf("apps/%s/firewall/ingressRules/%s", config["project"].(string), config["priority"].(string))
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
		// google_app_engine_flexible_app_version or google_app_engine_standard_app_version
	case res_type == "google_app_engine_flexible_app_version" || res_type == "google_app_engine_standard_app_version":
		if config["project"] != nil && config["service"] != nil && config["version_id"] != nil {
			id := fmt.Sprintf("apps/%s/services/%s/versions/%s", config["project"].(string), config["service"].(string), config["version_id"].(string))
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
		// google_app_engine_service_network_settings or google_app_engine_service_split_traffic
	case res_type == "google_app_engine_service_network_settings" || res_type == "google_app_engine_service_split_traffic":
		if config["project"] != nil && config["service"] != nil {
			id := fmt.Sprintf("apps/%s/services/%s", config["project"].(string), config["service"].(string))
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
		// google_artifact_registry_repository or google_artifact_registry_repository_iam_policy
	case res_type == "google_artifact_registry_repository" || res_type == "google_artifact_registry_repository_iam_policy":
		if config["project"] != nil && config["location"] != nil && config["repository_id"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", config["project"].(string), config["location"].(string), config["repository_id"].(string))
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
		// google_artifact_registry_repository_iam_member
	case res_type == "google_artifact_registry_repository_iam_member":
		if config["project"] != nil && config["location"] != nil && config["repository_id"] != nil && config["role"] != nil && config["member"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/repositories/%s %s %s", config["project"].(string), config["location"].(string), config["repository_id"].(string), config["role"].(string), config["member"].(string))
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
		// google_artifact_registry_repository_iam_binding
	case res_type == "google_artifact_registry_repository_iam_binding":
		if config["project"] != nil && config["location"] != nil && config["repository_id"] != nil && config["role"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/repositories/%s %s", config["project"].(string), config["location"].(string), config["repository_id"].(string), config["role"].(string))
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
		// google_artifact_registry_vpcsc_config
	case res_type == "google_artifact_registry_vpcsc_config":
		if config["project"] != nil && config["location"] != nil && config["name"] != nil {
			id := fmt.Sprintf("projects/%s/locations/%s/vpcscConfig/%s", config["project"].(string), config["location"].(string), config["name"].(string))
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
		// google_compute_instance
	case res_type == "google_compute_instance":
		if config["project"] != nil && config["zone"] != nil && config["name"] != nil {
			id := fmt.Sprintf("projects/%s/zones/%s/instances/%s", config["project"].(string), config["zone"].(string), config["name"].(string))
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
		// google_service_account
	case res_type == "google_service_account":
		if config["project"] != nil && config["account_id"] != nil {
			id := fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", config["project"].(string), config["account_id"].(string), config["project"].(string))
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
