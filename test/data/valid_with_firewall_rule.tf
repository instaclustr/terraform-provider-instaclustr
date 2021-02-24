provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
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
    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "apache-cassandra-3.11.8.ic2"
        options = {
            auth_n_authz = true
            use_private_broadcast_rpc_address = false
            client_encryption = false
            lucene_enabled = false
            continuous_backup_enabled = true
        }
    }
}

resource "instaclustr_firewall_rule" "valid_with_firewall_rule" {
    cluster_id = "${instaclustr_cluster.valid_with_firewall_rule.id}"
    rule_cidr = "10.220.0.0/21"
    rules = [
        {
            type = "CASSANDRA"
        }
    ]
}
