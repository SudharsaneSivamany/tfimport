module "vpc" {
  source = "terraform-google-modules/network/google"

  for_each = { for entry in var.vpc : entry.vpc_name => entry }

  project_id   = each.value.project_id
  network_name = each.value.vpc_name
  subnets      = each.value.subnets
}

variable "vpc" {
  type = list(object({
    project_id = string
    vpc_name   = string
    subnets = list(object({
      subnet_name   = string
      subnet_ip     = string
      subnet_region = string
    }))
  }))
  default = [{
    project_id = "<UPDATE PROJECT ID>" #UPDATE HERE 
    vpc_name   = "manual-vpc-1"
    subnets = [{
      subnet_name   = "manual-subnet-1"
      subnet_ip     = "10.0.0.0/24"
      subnet_region = "us-central1"
      },
      {
        subnet_name   = "tf-subnet-2"
        subnet_ip     = "10.0.1.0/24"
        subnet_region = "us-central1"
      }
    ]
  }]
}
