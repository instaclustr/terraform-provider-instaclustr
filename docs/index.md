---
page_title: "instaclustr Provider"
subcategory: ""
description: |-
  
---

# Instaclustr Provider

A [Terraform](http://terraform.io) provider for managing Instaclustr Platform resources.  

It provides a flexible set of resources for provisioning and managing [Instaclustr based clusters](http://instaclustr.com/) via the use of Terraform.  

For installation instructions and source code visit the [Github project page](https://github.com/instaclustr/terraform-provider-instaclustr)

For further information about Instaclustr, please see [FAQ](https://www.instaclustr.com/resources/faqs/) and [Support](https://support.instaclustr.com/hc/en-us) 

Use the navigation to the left to read about the available resources.


### Example Usage

```
terraform {
  required_providers {
    instaclustr = {
      source = "instaclustr/instaclustr"
      version = ">= 1.0.0, < 2.0.0"
    }
  }
}

variable "api_key" {
 type = string
 default = "xxx"
}

provider "instaclustr" {
    username= "<Your instaclustr username>"
    api_key = var.api_key
}

resource "instaclustr_cluster" "example" {
  cluster_name = "testcluster"
  node_size = "t3.small"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "AWS_VPC",
    tags = {
      "myTag" = "myTagValue"
    }
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.8"
    options = {
      auth_n_authz = true
    }
  }

  bundle {
    bundle = "SPARK"
    version = "2.3.2"
  }
}
```
