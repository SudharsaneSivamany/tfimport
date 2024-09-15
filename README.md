# tfimport

tfimport tool is for creating import blocks to import state of your existing infrastructure into your terraform modules or resources address in tfstate.

# Arguements
```bash
./tfimport setup   #to generate provider.json 

./tfimport plan    #to generate import.tf
```

## Providers Supported
1. Google Cloud Platform
    * google
    * google-beta
    * googleworkspace
2. Microsoft Azure
    * azurerm

## Pre-requisites
1. Terraform version >= 1.7.2
