---
page_title: "instaclustr_cluster Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource `instaclustr_cluster`

A resource for managing clusters on Instaclustr Managed Platform. A cluster contains a base application and several add-ons.

## Schema

Property | Description | Default
---------|-------------|--------
`cluster_name`|The name of new cluster. May contain a combination of letters, numbers and underscores with a maximum length of 32 characters.|Required
`node_size`|Desired node size. See [here](https://developer.instaclustr.com/#operation/extendedProvisionRequestHandler) for more details. This field is updatable.|Required
`data_centre`|Desired data centre. See [here](https://developer.instaclustr.com/#operation/extendedProvisionRequestHandler) for more details.|Required
`sla_tier`|Accepts PRODUCTION/NON_PRODUCTION. The SLA Tier feature on the Instaclustr console is used to classify clusters as either production and non_production. See [here](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more details.|NON_PRODUCTION
`cluster_network`|The private network address block for the cluster specified using CIDR address notation. The network must have a prefix length between /12 and /22 and must be part of a private address space.|10.224.0.0/12
`private_network_cluster`|Accepts true/false. Creates the cluster with private network only.|false
`pci_compliant_cluster`|Accepts true/false. Creates the cluster with PCI compliance enabled.|false
`cluster_provider`|The information of infrastructure provider. See below for its properties.|Required
`rack_allocation`|The number of resources to use. See below for its properties.|Optional, but Required for some Bundle types.
`bundle`|Array of bundle information. See below for its properties.|Required
`kafka_rest_proxy_user_password`|The password of kafka rest proxy bundle user, if it is a Kafka cluster with rest-proxy addon. This field is updatable and requires `wait_for_state` to be `RUNNING`.|Optional
`kafka_schema_registry_user_password`|The password of kafka schema registry bundle user, if it is a Kafka cluster with schema-registry addon. This field is updatable and requires `wait_for_state` to be `RUNNING`.|Optional
`wait_for_state`|The expected state of the cluster before completing the resource creation. Skipping this field will asynchronously create the cluster.|Optional (valid states are RUNNING and PROVISIONED)


### cluster_provider

Property | Description | Default
---------|-------------|--------
`name`|The name of the Cluster Provider. Accepts AWS_VPC, AZURE, and GCP.|Required
`account_name`|For customers running in their own account. Your provider account can be found on the ‘Account’ tab on the console, or the “Provider Account” property on any existing cluster.|""
`tags`|If specified, the value is a map from tag key to value. For restrictions, refer to the [AWS User Guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-restrictions). Tags are defined per cluster and will be applied to every instance in the cluster.|""
`resource_group`|AZURE only, if specified, the value is name for an Azure Resource Group which the resources will be provisioned into.|""
`disk_encryption_key`|Specify an KMS encryption key to encrypt data on nodes. KMS encryption key must be set in Account settings before provisioning an encrypted cluster.|""
`custom_virtual_network_id`|Specify a custom AWS VPC ID to use for customers provisioning in their own account. <b><i>Note:</i></b> Using this option requires that the cluster_network match the IPv4 CIDR block of the specified VPC ID.|""

`rack_allocation`

Property | Description | Default
---------|-------------|--------
`number_of_racks`|Number of racks to use when allocating nodes. Max allowed is 5|Required
`nodes_per_rack`|Number of nodes per rack. Max allowed is 10|Required

`bundle`

Property | Description | Default
---------|-------------|--------
`bundle`| See [Bundles and Versions below](#bundles-and-versions)|Required
`version`| See [Bundles and Versions below](#bundles-and-versions)|Required
`options`|Options and add-ons for the given bundle. See `bundle.options` below for its properties|{} (empty)

## Bundles and Versions

Bundle | Versions | Compatible With
---------|-------------|---------------
APACHE_CASSANDRA|2.2.18, 3.0.19, 3.11.8, 4.0 (preview)|
SPARK|2.1.3, 2.3.2|APACHE_CASSANDRA
KAFKA|2.1.1, 2.3.1, 2.4.1, 2.5.1, 2.6.1|
KAFKA_REST_PROXY|5.0.0|KAFKA
KAFKA_SCHEMA_REGISTRY|5.0.0|KAFKA
ELASTICSEARCH|opendistro-for-elasticsearch:1.4.0
KAFKA_CONNECT|2.3.1, 2.4.1, 2.5.1, 2.6.1|
REDIS|6.0.9|
APACHE_ZOOKEEPER|3.5.8|
POSTGRESQL|13.4|

`bundle.options`

Property | Description | For Bundles | Default
---------|-------------|-------------|--------
`auth_n_authz`|Accepts true/false. Enables Password Authentication and User Authorization.|Cassandra|false
`client_encryption`|Accepts true/false. Enables Client ⇄ Node Encryption.|Cassandra, Kafka, Elasticsearch, Spark, Redis, ZooKeeper, PostgreSQL|false
`dedicated_master_nodes`|Accepts true/false. Enables Dedicated Master Nodes.|Elasticsearch|false
`master_node_size`|Desired master node size. See [here](https://developer.instaclustr.com/#operation/extendedProvisionRequestHandler) for more details.|Elasticsearch|Required
`security_plugin`|Accepts true/false. Enables Security Plugin. This option gives an extra layer of security to the cluster. This will automatically enable internode encryption.|Elasticsearch|false
`use_private_broadcast_rpc_address`|Accepts true/false. Enables broadcast of private IPs for auto-discovery.|Cassandra|false
`lucene_enabled`|Accepts true/false. Enabled Cassandra Lucene Index Plugin.|Cassandra|false
`continuous_backup_enabled`|Accepts true/false. Enables commitlog backups and increases the frequency of the default snapshot backups.|Cassandra|false
`number_partitions`|Default number of partitions to be assigned per topic.|Kafka|Number of nodes
`auto_create_topics`|Accepts true/false. Enable to allow brokers to automatically create a topic, if it does not already exist, when messages are published to it.|Kafka|true
`delete_topics`|Accepts true/false. Enable to allow topics to be deleted via the `ic-kafka-topics` tool.|Kafka|true
`password_authentication`|Accepts true/false. Require clients to provide credentials — a username & API Key — to connect to the Spark Jobserver.|Spark|false
`target_kafka_cluster_id`|GUID of the Instaclustr managed Kafka Cluster Id you wish to connect to. Must be in the same Instaclustr account.|Kafka Connect|Required
`vpc_type`|Available options: `KAFKA_VPC`, `VPC_PEERED`, `SEPARATE_VPC`. Only required if targeting an Instaclustr managed cluster.|Kafka Connect|`SEPARATE_VPC`
`aws_access_key`, `aws_secret_key`, `s3_bucket_name`|Access information for the S3 bucket where you will store your custom connectors. (if using AWS)|Kafka Connect
`azure_storage_account_name`, `azure_storage_account_key`, `azure_storage_container_name`|Access information for the Azure Storage container where you will store your custom connectors.|Kafka Connect
`ssl_enabled_protocols`, `ssl_truststore_password`, `ssl_protocol`, `security_protocol`, `sasl_mechanism`, `sasl_jaas_config`, `bootstrap_servers`|Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster|Kafka Connect
`truststore`|Base64 encoded version of the TLS trust store (in JKS format) used to connect to your Kafka Cluster. Only required if connecting to a Non-Instaclustr managed Kafka Cluster with TLS enabled|Kafka Connect
`master_nodes`|The number of Master nodes in a generated Redis Cluster.|Redis|Required (Integers)
`replica_nodes`|The number of Replica nodes in a generated Redis Cluster.|Redis|Required (Integers)
`password_auth`|Accepts true/false. Enables Password Authentication and User Authorization.|Redis|false
`dedicated_zookeeper`|Indicate whether this Kafka cluster should allocate dedicated Zookeeper nodes|Kafka|false
`zookeeper_node_size`|If `dedicated_zookeeper` is true, then it is the node size for the dedicated Zookeeper nodes. Have a look [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-create-cluster) (Kafka bundle options table) for node size options. |Kafka
`zookeeper_node_count`|Indicates how many nodes are allocated to be Zookeeper nodes. For KAFKA bundle, if `dedicated_zookeeper` is false, then it indicates how many Kafka nodes also have ZooKeeper services in them. |Kafka, ZooKeeper
`postgresql_node_count`|The number of nodes in a generated PostgreSQL cluster.|Postgresql|Required (Integers)


