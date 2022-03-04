# required for terraform version >= 13.
# Remove the terraform block if you're using a terraform version lower
terraform {
  required_providers {
    instaclustr = {
//      source = "instaclustr/instaclustr"
      //Change the source as per below to work with a local development copy on terraform version >=13
      source = "terraform.instaclustr.com/instaclustr/instaclustr"
      version = ">= 1.0.0"
    }
  }
}

provider "instaclustr" {
  username = "alwyn"
  api_key = "ae27289d3e155df440b4a5df2acdcba4"
  api_hostname = "http://localhost:8090"
}

resource "instaclustr_cluster" "alwyn-terraform-postgresql" {
  cluster_name = "tharindu-terraform-postgresql"
  node_size = "PGS-DEV-t3.small-5"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    nodes_per_rack = 1
    number_of_racks = 1
  }

  lifecycle {
    ignore_changes = [rack_allocation]
  }

  bundle {
    bundle = "POSTGRESQL"
    version = "14.1"
    options = {
      postgresql_node_count = 2,
      client_encryption = true,
      replication_mode = "SYNCHRONOUS",
      synchronous_mode_strict = true
    }
  }
}