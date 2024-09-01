# Steps 

On Microsoft Azure Portal, navigate to management groups and create managemnt groups in the below order under root tenant group
Tenant root
|__mg1
   |__mg2
      |__mg2

```bash
az login
```

Upgrade Terrform to latest version

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform

```

```bash
wget https://github.com/SudharsaneSivamany/tfimport/releases/download/v0.1.0/tfimport_0.1.0_linux_amd64.tar.gz

tar -xzvf tfimport_0.1.0_linux_amd64.tar.gz 

cp tfimport example/azure/

terraform init

terraform plan   
```

You will see 3 resource are planned for create. Also Note if terraform plan ends up with error, it has to be fixed before tfimport is executed 

```bash
./tfimport setup

```

tfimport setup will create provider.json. Please update the subsription id for the providers


```bash
./tfimport plan
```


After this you will see import.tf is getting updated like below 
```hcl

import {
  id = "/providers/Microsoft.Management/managementGroups/mg1"
  to = module.management-group.azurerm_management_group.level_1_group["mg1"]
}


import {
  id = "/providers/Microsoft.Management/managementGroups/mg2"
  to = module.management-group.azurerm_management_group.level_2_group["mg2"]
}


import {
  id = "/providers/Microsoft.Management/managementGroups/mg3"
  to = module.management-group.azurerm_management_group.level_3_group["mg3"]
}


```

```bash
terraform apply


terraform destroy
```




