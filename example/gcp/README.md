# Steps 

PLease login to google cloud shell

```bash
gcloud config set project <project-id>

gcloud compute networks create manual-vpc-1 --subnet-mode=custom --bgp-routing-mode=global

gcloud compute networks subnets create manual-subnet-1 --network=manual-vpc-1 --range=10.0.0.0/24 --region=us-central1
```

Upgrade Terrform to latest version

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform

```

Update the project-id in the main.tf

```bash
cd ../../

go build . 

cp tfimport example/gcp/

terraform init

terraform plan   
```

You will see 3 resource are planned for create. 

Note: if terraform plan ends up with error, it has to be fixed before tfimport is executed 

```bash
./tfimport setup

./tfimport plan
```


After this you will see import.tf is getting updated like below 
```hcl
import {
  id = "projects/<your project id>/regions/us-central1/subnetworks/manual-subnet"
  to = module.vpc["manual-vpc-1"].module.subnets.google_compute_subnetwork.subnetwork["us-central1/manual-subnet-1"]
}



import {
  id = "projects/<your project id>/global/networks/manual-vpc-1"
  to = module.vpc["manual-vpc-1"].module.vpc.google_compute_network.network
}


```

```bash
terraform apply


terraform destroy
```




