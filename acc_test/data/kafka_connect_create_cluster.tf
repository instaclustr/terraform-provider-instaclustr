# kafka connect acceptance test. Using the same base kafka cluster as kafka user testing to minimize cluster creation
resource "instaclustr_firewall_rule" "valid_to_kafka_cluster" {
    cluster_id = instaclustr_cluster.kafka_cluster.id
    rule_cidr = "0.0.0.0/0"
    rules = [{type = "KAFKA"}]
}

resource "instaclustr_kafka_user" "kafka_user_kc_charlie" {
  cluster_id = instaclustr_cluster.kafka_cluster.id
  username = "Charlie"
  password = "Charlie123!"
  initial_permissions = "standard"
  override_existing_user = false
}

resource "instaclustr_cluster" "validKCAws" {
    wait_for_state = "RUNNING"
    cluster_name = "tf_testacc_kc_aws_connectors"
    node_size = "KCN-DEV-t4g.medium-30"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.128.0/18"
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
        version = "%s"
        options = {
            target_kafka_cluster_id = instaclustr_cluster.kafka_cluster.id
            vpc_type = "VPC_PEERED"
            s3_bucket_name = "%s"
        }
    }
  
    kafka_connect_credential  {
        aws_access_key = "%s"
        aws_secret_key = "%s"
    }
   
}

resource "instaclustr_cluster" "validKCAzure" {
    wait_for_state = "RUNNING"
    cluster_name = "tf_testacc_kc_azure_connectors"
    node_size = "Standard_D2s_v3-10"
    data_centre = "CENTRAL_US"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.192.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false
    cluster_provider = {
        name = "AZURE_AZ"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }

    kafka_connect_credential  {
        azure_storage_account_name = "%s"
        azure_storage_account_key = "%s"
        sasl_jaas_config = "org.apache.kafka.common.security.scram.ScramLoginModule required username=\"Charlie\" password=\"Charlie123!\";"
    }

    bundle {
        bundle = "KAFKA_CONNECT"
        version = "%s"
        options = {
            vpc_type = "SEPARATE_VPC"
            azure_storage_container_name = "%s"
            security_protocol = "SASL_PLAINTEXT"
            sasl_mechanism = "SCRAM-SHA-256"
            bootstrap_servers = format("%%s:9092", join(":9092,",instaclustr_cluster.kafka_cluster.public_contact_point))
        }
    }
}
