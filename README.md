# tfimport

tfimport tool is for creating import blocks to import state of your existing infrastructure into your terraform modules or resources.

# Arguements
```bash
./tfimport setup   #to generate provider.json 

./tfimport plan    #to generate import.tf
```

## Providers Supported
1. Google Cloud Platform
2. Microsoft Azure

## Pre-requisites
1. Terraform version >= 1.7.2
2. Golang package install

## Google Cloud Platform supported resource blocks:
* google_compute_network
* google_compute_subnetwork
* google_api_gateway_api
* google_api_gateway_api_config
* google_api_gateway_api_config_iam_member
* google_api_gateway_api_config_iam_binding
* google_api_gateway_api_config_iam_policy
* google_api_gateway_api_iam_member
* google_api_gateway_api_iam_binding
* google_api_gateway_api_iam_policy
* google_api_gateway_gateway
* google_api_gateway_gateway_iam_member
* google_api_gateway_gateway_iam_binding
* google_api_gateway_gateway_iam_policy
* google_alloydb_backup
* google_alloydb_cluster
* google_alloydb_instance
* google_alloydb_user
* google_app_engine_application
* google_app_engine_application_url_dispatch_rules
* google_app_engine_domain_mapping
* google_app_engine_firewall_rule
* google_app_engine_flexible_app_version
* google_app_engine_standard_app_version
* google_app_engine_service_network_settings
* google_app_engine_service_split_traffic
* google_artifact_registry_repository
* google_artifact_registry_repository_iam_policy
* google_artifact_registry_repository_iam_member
* google_artifact_registry_repository_iam_binding
* google_artifact_registry_vpcsc_config
* google_compute_instance
* google_service_account

## Microsoft Azure supported resource blocks:
* azurerm_resource_group
* azurerm_virtual_network
* azurerm_subnet
* azurerm_route_table
* azurerm_network_security_group
* azurerm_network_security_rule
* azurerm_virtual_network_peering
* azurerm_subnet_route_table_association
* azurerm_subnet_network_security_group_association
* azurerm_management_group
* azurerm_storage_account
* azurerm_virtual_machine
* azurerm_linux_virtual_machine
* azurerm_windows_virtual_machine
* azurerm_virtual_desktop_host_pool
* azurerm_virtual_desktop_application_group
* azurerm_virtual_desktop_workspace
* azurerm_virtual_desktop_workspace_application_group_association
* azurerm_virtual_desktop_scaling_plan
* azurerm_virtual_desktop_application