provider "instaclustr" {
    username = "<Your instaclustr username here>"
    api_key = "<Your provisioning API key here>"
}

resource "instaclustr_cluster" "example2" {
    cluster_name = "testcluster"
    node_size = "t3.small"
    data_centre = "US_WEST_2"
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
    bundles = [
        {
            bundle = "APACHE_CASSANDRA"
            version = "3.11.4"
        },
        {
            bundle = "SPARK"
            version = "apache-spark:2.3.2"
        },
        {
            bundle = "ZEPPELIN"
            version = "apache-zeppelin:0.8.0-spark-2.3.2"
        }
    ]
}

resource "instaclustr_firewall_rule" "test3" {
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
