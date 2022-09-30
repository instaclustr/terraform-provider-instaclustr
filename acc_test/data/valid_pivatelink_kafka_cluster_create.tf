provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "validPrivateLinkKafka" {
  cluster_name = "example_kafka_privatelink"
  node_size = "KFK-PRD-r6g.large-250"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  private_network_cluster = true
  cluster_network = "192.168.0.0/18"
  wait_for_state = "RUNNING"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "KAFKA"
    version = "3.1.2"
    options = {
      auto_create_topics = true
      client_encryption = false
      delete_topics = true
      number_partitions = 3
      advertised_host_name = "kafka.test.com"
      zookeeper_node_count = 3
    }
  }

  private_link {
    iam_principal_arns = ["arn:aws:iam::123456789012:root"]
  }
}
