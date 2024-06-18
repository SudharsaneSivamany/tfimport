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