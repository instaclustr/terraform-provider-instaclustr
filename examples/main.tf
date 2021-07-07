# required for terraform version >= 13.
# Remove the terraform block if you're using a terraform version lower
terraform {
  required_providers {
    instaclustr = {
      source = "instaclustr/instaclustr"
      //Change the source as per below to work with a local development copy on terraform version >=13
      //source = "terraform.instaclustr.com/instaclustr/instaclustr"
      version = ">= 1.0.0"
    }
  }
}

provider "instaclustr" {
  username = "<Your instaclustr username here>"
  api_key = "<Your provisioning API key here>"
}

resource "instaclustr_encryption_key" "add_ebs_key" {
  alias = "testkey"
  arn = "<Your KMS key ARN here>"
  provider = "instaclustr"
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
}

data "instaclustr_cluster_credentials" "example_credentials" {
  cluster_id = "${instaclustr_cluster.example.id}"
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
    version = "3.11.8"
  }
}

resource "instaclustr_firewall_rule" "example_firewall_rule" {
  cluster_id = "${instaclustr_cluster.example.id}"
  rule_cidr = "10.1.0.0/16"
  rules = [
    {
      type = "CASSANDRA"
    }
  ]
}

resource "instaclustr_firewall_rule" "example_firewall_rule_sg" {
  cluster_id = "${instaclustr_cluster.example.id}"
  rule_security_group_id = "sg-0123abcde456ffabc"
  rules = [
    {
      type = "CASSANDRA"
    }
  ]
}

resource "instaclustr_vpc_peering" "example_vpc_peering" {
  cluster_id = "${instaclustr_cluster.example.id}"
  peer_vpc_id = "vpc-123456"
  peer_account_id = "1234567890"
  peer_subnet = "10.0.0.0/20"
}

// Updating the kafka-schema-registry and the kafka-rest-proxy bundle user passwords at the cluster creation time
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
    version = "2.5.1"
    options = {
      auth_n_authz = true
      dedicated_zookeeper = true
      zookeeper_node_size = "zk-production-m5.large-60"
      zookeeper_node_count = 3
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
  kafka_rest_proxy_user_password = "RestProxyTest123test!" // new password for rest proxy bundle user
  kafka_schema_registry_user_password = "SchemaRegistryTest123test!" // new password for schema registry bundle user

  wait_for_state = "RUNNING" // the required state of the cluster before doing the bundle user password updates
}

resource "instaclustr_cluster" "example-elasticsearch" {
  cluster_name = "es-cluster"
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
      data_node_size = "m5l-250-v2",
      kibana_node_size = "m5l-250-v2",
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
      target_kafka_cluster_id = "${instaclustr_cluster.example_kafka.id}"
      vpc_id = "SEPARATE_VPC"
    }
  }
}

resource "instaclustr_kafka_user" "kafka_user_charlie" {
  cluster_id = "${instaclustr_cluster.example_kafka.id}"
  username = "charlie"
  password = "charlie123!"
}

resource "instaclustr_kafka_user" "kafka_user_harley" {
  cluster_id = "${instaclustr_cluster.example_kafka.id}"
  username = "harley"
  password = "harley123!"
  initial_permissions = "standard"
  sasl_scram_mechanism = "SCRAM-SHA-512"
}

data "instaclustr_kafka_user_list" "kafka_user_list" {
  cluster_id = "${instaclustr_cluster.example_kafka.id}"
}

resource "instaclustr_cluster" "private_cluster_example" {
  cluster_name = "testcluster"
  node_size = "m5l-250-v2"
  data_centre = "US_EAST_1"
  sla_tier = "PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = true
  cluster_provider = {
    name = "AWS_VPC",
  }
  rack_allocation = {
    number_of_racks = 2
    nodes_per_rack = 1
  }
  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.8"
  }
}

resource "instaclustr_cluster" "example-redis" {
  cluster_name = "testcluster"
  node_size = "t3.small-20-r"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }

  bundle {
    bundle = "REDIS"
    version = "6.0.4"
    options = {
      master_nodes = 3,
      replica_nodes = 3
    }
  }
}
