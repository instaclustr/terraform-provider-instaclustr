terraform {
  required_providers {
    instaclustr = {
      source = "terraform.instaclustr.com/instaclustr/instaclustr"
      version = ">= 2.0.0, < 3.0.0"
    }
  }
}

provider "instaclustr" {
  terraform_key = "Instaclustr-Terraform Shami:8226f255aef347152881557cd521ad2b"
}
data "instaclustr_redis_cluster_v2_instance" "example" {
 id = "1f948192-70d3-46d1-be35-231abddedfca"
}

output "example" {
  value = data.instaclustr_redis_cluster_v2_instance.example
}


