provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "valid_with_firewall_rule" {
    cluster_name = "tf-provider-firewall-rule-test"
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
        }
    ]
}

resource "instaclustr_firewall_rule" "valid_with_firewall_rule" {
    cluster_id = "${instaclustr_cluster.valid_with_firewall_rule.id}"
    rule_cidr = "10.1.0.0/16"
    rules = [
        { 
            type = "CASSANDRA"
        }
    ]
}