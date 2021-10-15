provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
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
        version = "2.7.1"
        options = {
            target_kafka_cluster_id = "%s"
            vpc_id = "SEPARATE_VPC"
            aws_access_key = "%s"
            aws_secret_key = "%s"
            s3_bucket_name = "%s"
        }
    }
}
