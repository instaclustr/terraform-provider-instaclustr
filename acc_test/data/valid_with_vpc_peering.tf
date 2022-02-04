provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "valid_with_vpc_peering" {
    cluster_name = "tf_provider_vpc_peering_test"
    node_size = "t3.small-v2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    wait_for_state = "RUNNING"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.8"
        options = {
            auth_n_authz = false
            client_encryption = false
            continuous_backup_enabled = false
            lucene_enabled = false
            use_private_broadcast_rpc_address = false
        }
    }
}

resource "instaclustr_vpc_peering" "valid_with_vpc_peering" {
    cluster_id = "${instaclustr_cluster.valid_with_vpc_peering.cluster_id}"
    
    peer_vpc_id = "vpc-12345678"
    peer_account_id = "494111121110"
    peer_subnets = toset(["10.128.176.0/20", "10.129.176.0/20"])
}

resource "instaclustr_vpc_peering" "valid_with_vpc_peering_single_subnet" {
    cluster_id = "${instaclustr_cluster.valid_with_vpc_peering.cluster_id}"
    peer_vpc_id = "vpc-12345679"
    peer_account_id = "494111121110"
    peer_subnets = toset(["10.130.176.0/20"])
}

resource "instaclustr_vpc_peering" "valid_with_vpc_peering_legacy" {
    cluster_id = "${instaclustr_cluster.valid_with_vpc_peering.cluster_id}"
    peer_vpc_id = "vpc-12345680"
    peer_account_id = "494111121110"
    peer_subnet = "10.131.176.0/20"
}
resource "instaclustr_cluster" "gcp_example" {
  cluster_name = "testclustergcp"
  node_size = "n1-standard-2"
  data_centre = "us-east1"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  wait_for_state = "RUNNING"
  private_network_cluster = false
  cluster_provider = {
    name = "GCP"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "apache-cassandra-3.11.8.ic4"
    options = {
      auth_n_authz = true
    }
  }
}
resource "instaclustr_vpc_peering_gcp" "gcp_example" {
  cluster_id = "${instaclustr_cluster.gcp_example.id}"
  name="name"
  peer_vpc_network_name = "my-vpc1"
  peer_project_id = "instaclustr-dev"
  peer_subnets = toset(["10.10.0.0/16", "10.11.0.0/16"])
  
}
