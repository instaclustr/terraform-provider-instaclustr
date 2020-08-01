provider "instaclustr" {
  username = "<Your instaclustr username here>"
  api_key = "<Your provisioning API key here>"
}

resource "instaclustr_encryption_key" "add_ebs_key" {
    alias = "testkey"
    arn = "<Your KMS key ARN here>"
}

resource "instaclustr_cluster" "example2" {
  cluster_name = "testcluster"
  node_size = "t3.small"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "AWS_VPC",
    disk_encryption_key = "${instaclustr_encryption_key.add_ebs_key.key_id}"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.4"
    options = {
      auth_n_authz = true
    }
  }

  bundle {
    bundle = "SPARK"
    version = "apache-spark:2.3.2"
  }

  bundle {
    bundle = "ZEPPELIN"
    version = "apache-zeppelin:0.8.0-spark-2.3.2"
  }
}

resource "instaclustr_cluster" "custom_vpc_example" {
  cluster_name = "testcluster"
  node_size = "t3.small"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "10.100.0.0/16"
  private_network_cluster = false
  cluster_provider = {
    name = "AWS_VPC",
    account_name = "My RIYOA Account"
    custom_virtual_network_id = "vpc-0a1b2c3d4e5f6g7h8"
  }
  rack_allocation = {
    number_of_racks = 2
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.4"
  }
}

resource "instaclustr_firewall_rule" "example_firewall_rule" {
  cluster_id = "${instaclustr_cluster.example2.id}"
  rule_cidr = "10.1.0.0/16"
  rules = [
    {
      type = "CASSANDRA"
    }
  ]
}

resource "instaclustr_vpc_peering" "example_vpc_peering" {
  cluster_id = "${instaclustr_cluster.example.cluster_id}"
  peer_vpc_id = "vpc-123456"
  peer_account_id = "1234567890"
  peer_subnet = "10.0.0.0/20"
}

resource "instaclustr_cluster" "example_kafka" {
  cluster_name = "test_kafka"
  node_size = "r5.large-500-gp2"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "KAFKA"
    version = "2.3.1"
    options = {
      auth_n_authz = true
    }
  }

  bundle {
    bundle = "KAFKA_REST_PROXY"
    version = "5.0.0"
  }

  bundle {
    bundle = "KAFKA_SCHEMA_REGISTRY"
    version = "5.0.0"
  }
}

resource "instaclustr_cluster" "example-elasticsearch" {
  cluster_name = "es-cluster"
  node_size = "m5l-250-v2"
  data_centre = "US_EAST_1"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "ELASTICSEARCH"
    version = "opendistro-for-elasticsearch:1.4.0"
    options = {
      client_encryption = true,
      dedicated_master_nodes = true,
      master_node_size = "m5l-250-v2",
      security_plugin = true
    }
  }
}

resource "instaclustr_cluster" "validKC" {
    cluster_name = "testcluster"
    node_size = "t3.medium-10-gp2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }

    bundle {
        bundle = "KAFKA_CONNECT"
        version = "2.3.1"
        options = {
            target_kafka_cluster_id = "${instaclustr_cluster.example_kafka.cluster_id}"
            vpc_id = "SEPARATE_VPC"
        }
    }
}

resource "instaclustr_kafka_user" "kafka_user_charlie" {
  cluster_id = "${instaclustr_clustr.example_kafka.cluster_id}"
  username = "charlie"
  password = "charlie123!"
  initial_permissions = "none"
}                       
                              
resource "instaclustr_kafka_user_list" "kafka_user_list" { 
  cluster_id = "${instaclustr_clustr.example_kafka.cluster_id}"
}  
