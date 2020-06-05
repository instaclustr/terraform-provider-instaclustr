# Terraform Instaclustr Provider

A [Terraform](http://terraform.io) provider for managing Instaclustr Platform resources.  

It provides a flexible set of resources for provisioning and managing [Instaclustr based clusters](http://instaclustr.com/) via the use of Terraform.  

For general information about Terraform, visit the [official website](https://terraform.io/) and [GitHub project page](https://github.com/hashicorp/terraform).

For further information about Instaclustr, please see [FAQ](https://www.instaclustr.com/resources/faqs/) and [Support](https://support.instaclustr.com/hc/en-us) 

## Key benefits

- Removes the need to write custom code integration directly with the Instaclustr API
- Instaclustr based infrastructure as code deployments with minimal effort
- Ease of integration into existing terraform or automated CI/CD based workflows
- Ease of customisation and configuration in order to meet operational requirements
- Use of existing Instaclustr authentication methodologies

## Requirements

- Terraform 0.10.x or higher
- Go 1.14 or higher

## Building The Provider

Clone the provider source code

```sh
$ mkdir -p $GOPATH/src/github.com/instaclustr; cd $GOPATH/src/github.com/instaclustr
$ git clone https://github.com/instaclustr/terraform-provider-instaclustr.git
```

Build the source into executable binary

```sh
$ cd $GOPATH/src/github.com/instaclustr/terraform-provider-instaclustr
$ make build
```

Install the provider

```sh
$ cd $GOPATH/src/github.com/instaclustr/terraform-provider-instaclustr
$ make install
```

For further details on Provider installation please see the [Terraform installation guide](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin).

## Authentication

This provider requires an API Key in order to provision Instaclustr resources. To create an API key, please log into the [Instaclustr Console](https://console.instaclustr.com) or signup for an account [here](https://console.instaclustr.com/user/signup) if you dont have one.  Navigate to `Account` -> `API Keys` page, locate the `Provisioning` role and click `Generate Key`.  This username and API key combination should then be placed into the provider configuration:

```
provider "instaclustr" {
    username = "<Your instaclustr username here>"
    api_key = "<Your provisioning API key here>"
}
```
or just input them when Terraform indicates you.

## Example Usage

```
resource "instaclustr_cluster" "example" {
    cluster_name = "testcluster"
    node_size = "t3.small"
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
```

## Configuration
### Resources
### Resource:  `instaclustr_cluster`  
A resource for managing clusters on Instaclustr Managed Platform. A cluster contains a base application and several add-ons.

#### Properties
Property | Description | Default
---------|-------------|--------
cluster_name|The name of new cluster. May contain a combination of letters, numbers and underscores with a maximum length of 32 characters.|Required
node_size|Desired node size. See [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-reference-data-data-centres-and-node-sizes) for more details.|Required
data_centre|Desired data centre. See [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-reference-data-data-centres-and-node-sizes) for more details.|Required
sla_tier|Accepts PRODUCTION/NON_PRODUCTION. The SLA Tier feature on the Instaclustr console is used to classify clusters as either production and non_production. See [here](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more details.|NON_PRODUCTION
cluster_network|The private network address block for the cluster specified using CIDR address notation. The network must have a prefix length between /12 and /22 and must be part of a private address space.|10.224.0.0/12
private_network_cluster|Accepts true/false. Creates the cluster with private network only.|false
pci_compliant_cluster|Accepts true/false. Creates the cluster with PCI compliance enabled.|false
cluster_provider|The information of infrastructure provider. See below for its properties.|Required
rack_allocation|The number of resources to use. See below for its properties.|Required
bundle|Array of bundle information. See below for its properties.|Required

`cluster_provider`

Property | Description | Default
---------|-------------|--------
name|Accepts AWS_VPC now. The new cluster will be deployed on Amazon Web Service.|Required
account_name|For customers running in their own account. Your provider account can be found on the ‘Account’ tab on the console, or the “Provider Account” property on any existing cluster.|""
tags|If specified, the value is a map from tag key to value. For restrictions, refer to the [AWS User Guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-restrictions). Tags are defined per cluster and will be applied to every instance in the cluster.|""
resource_group|AZURE only, if specified, the value is name for an Azure Resource Group which the resources will be provisioned into.|""
disk_encryption_key|Specify an KMS encryption key to encrypt data on nodes. KMS encryption key must be set in Account settings before provisioning an encrypted cluster.|""
custom_virtual_network_id|Specify a custom AWS VPC ID to use for customers provisioning in their own account. <b><i>Note:</i></b> Using this option requires that the cluster_network match the IPv4 CIDR block of the specified VPC ID.|""

`rack_allocation`

Property | Description | Default
---------|-------------|--------
number_of_racks|Number of racks to use when allocating nodes. Max allowed is 5|Required
nodes_per_rack|Number of nodes per rack. Max allowed is 10|Required

`bundle`

Property | Description | Default
---------|-------------|--------
bundle| See [Bundles and Versions below](#bundles-and-versions)|Required
version| See [Bundles and Versions below](#bundles-and-versions)|Required
options|Options and add-ons for the given bundle. See `bundle.options` below for its properties|{} (empty)

`bundle.options` - _all properties listed are optional_

Property | Description | For Bundles | Default
---------|-------------|-------------|--------
auth_n_authz|Accepts true/false. Enables Password Authentication and User Authorization.|Cassandra|false
client_encryption|Accepts true/false. Enables Client ⇄ Node Encryption.|Cassandra, Kafka, Elasticsearch, Spark|false
dedicated_master_nodes|Accepts true/false. Enables Dedicated Master Nodes.|Elasticsearch|false
master_node_size|Desired master node size. See [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-reference-data-data-centres-and-node-sizes) for more details.|Elasticsearch|Required
security_plugin|Accepts true/false. Enables Security Plugin. This option gives an extra layer of security to the cluster. This will automatically enable internode encryption.|Elasticsearch|false
use_private_broadcast_rpc_address|Accepts true/false. Enables broadcast of private IPs for auto-discovery.|Cassandra|false
lucene_enabled|Accepts true/false. Enabled Cassandra Lucene Index Plugin.|Cassandra|false
continuous_backup_enabled|Accepts true/false. Enables commitlog backups and increases the frequency of the default snapshot backups.|Cassandra|false
number_partitions|Default number of partitions to be assigned per topic.|Kafka|Number of nodes
auto_create_topics|Accepts true/false. Enable to allow brokers to automatically create a topic, if it does not already exist, when messages are published to it.|Kafka|true
delete_topics|Accepts true/false. Enable to allow topics to be deleted via the `ic-kafka-topics` tool.|Kafka|true
password_authentication|Accepts true/false. Require clients to provide credentials — a username & API Key — to connect to the Spark Jobserver.|Spark|false
target_kafka_cluster_id|GUID of the Instaclustr managed Kafka Cluster Id you wish to connect to. Must be in the same Instaclustr account.|Kafka Connect|Required
vpc_type|Available options: `KAFKA_VPC`, `VPC_PEERED`, `SEPARATE_VPC`. Only required if targeting an Instaclustr managed cluster.|Kafka Connect|`SEPARATE_VPC`
aws_access_key, aws_secret_key, s3_bucket_name|Access information for the S3 bucket where you will store your custom connectors. (if using AWS)|Kafka Connect
azure_storage_account_name, azure_storage_account_key, azure_storage_container_name|Access information for the Azure Storage container where you will store your custom connectors.|Kafka Connect
ssl_enabled_protocols, ssl_truststore_password, ssl_protocol, security_protocol, sasl_mechanism, sasl_jaas_config, bootstrap_servers|Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster|Kafka Connect
truststore|Base64 encoded version of the TLS trust store (in JKS format) used to connect to your Kafka Cluster. Only required if connecting to a Non-Instaclustr managed Kafka Cluster with TLS enabled|Kafka Connect

### Resource:  `instaclustr_firewall_rule`                             
A resource for managing cluster firewall rules on Instaclustr Managed Platform. A firewall rule allows access to your Instaclustr cluster.

#### Properties
Property | Description | Default
---------|-------------|--------
cluster_id|The ID of an existing Instaclustr managed cluster|Required
rule_cidr|The network to add to your cluster firewall rule. Must be a valid IPv4 CIDR|Required
rules|List of rule types that the specified network is allowed access to. See below for rule options.|Required

`rules`

Property | Description | Default
---------|-------------|--------
type|Accepts CASSANDRA, SPARK, SPARK_JOBSERVER|Required

#### Example
```
resource "instaclustr_firewall_rule" "example" {
    cluster_id = "${instaclustr_cluster.example.id}"
    rule_cidr = "10.1.0.0/16"
    rules = [
        { 
            type = "CASSANDRA"
        }
    ]
}
```

### Resource: `instaclustr_vpc_peering`  
A resource for managing VPC peering connections on Instaclustr Managed Platform. This is only avaliable for clusters hosted with the AWS provider.

#### Properties
Property | Description | Default
---------|-------------|--------
cluster_id|The ID of an existing Instaclustr managed cluster| Not Required if cdc_id provided
cdc_id|The ID of an existing Instaclustr managed cluster data centre|Not Required if cluster_id provided
peer_vpc_id|The ID of the VPC with which you are creating the VPC peering connection|Required
peer_account_id|The account ID of the owner of the accepter VPC|Required
peer_subnet|The subnet for the VPC|Required
peer_region| The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request. | Not Required


#### Example
```
resource "instaclustr_vpc_peering" "example_vpc_peering" {
    cluster_id = "${instaclustr_cluster.example.cluster_id}"
    peer_vpc_id = "vpc-123456"
    peer_account_id = "1234567890"
    peer_subnet = "10.0.0.0/20"
}
```

### Resource: `instaclustr_encryption_key`  
A resource for managing EBS encryption of nodes with KMS keys. This is only avaliable for clusters hosted with the AWS provider.

#### Properties
Property | Description | Default
---------|-------------|--------
key_id|Internal ID of the KMS encryption key. Can be found via GET to `https://api.instaclustr.com/provisioning/v1/encryption-keys`|""
alias|KMS key alias, a human-readibly identifier specified alongside your KMS ARN|""
arn|KMS ARN, identifier specifying provider, location and key in a ':' value seperated string|""

#### Example
```
resource "instaclustr_encryption_key" "example_encryption_key" {
    alias: "virginia 1"
    arn: "arn:aws:kms:us-east-1:123456789012:key12345678-1234-1234-1234-123456789abc"
}
```

## Bundles and Versions

Bundle | Versions | Compatible With
---------|-------------|---------------
APACHE_CASSANDRA|2.1.19, 2.2.13, 3.0.14, 3.0.17, 3.0.18, 3.11, 3.11.3, 3.11.4|
SPARK|apache-spark:2.1.3, apache-spark:2.1.3.ic1, apache-spark:2.3.2|APACHE_CASSANDRA
ZEPPELIN|apache-zeppelin:0.8.0-spark-2.3.2, apache-zeppelin:0.7.1-spark-2.1.1|APACHE_CASSANDRA
KAFKA|2.1.1, 2.3.1, 2.4.1|
KAFKA_REST_PROXY|5.0.0|KAFKA
KAFKA_SCHEMA_REGISTRY|5.0.0|KAFKA
ELASTICSEARCH|opendistro-for-elasticsearch:1.4.0
KAFKA_CONNECT|2.3.1, 2.4.1|

### Migrating from 0.0.1 &rarr; 1.0.0+
A schema change has been made from 0.0.1 which no longer supports the `bundles` argument and uses `bundle` blocks instead. This change can cause `terraform apply` to fail with a message that `bundles` has been removed and/or updating isn't supported. To resolve this -<br>
1. Change all usages of the `bundles` argument &rarr; `bundle` blocks (example under example/main.tf)
2. In the .tfstate files, replace all keys named `bundles` with `bundle` in resources under the Instaclustr provider.

## Contributing

Firstly thanks!  We value your time and will do our best to review the PR as soon as possible. 

1. [Install golang](https://golang.org/doc/install#install)
2. Clone repository to: $GOPATH/src/github.com/instaclustr/terraform-provider-instaclust
3. Build the provider by `$ make build`
4. Run the tests by `$ make test`
5. Set up all of the environmental variables listed [below](#acceptance-test-environment-variables) to prepare for acceptance testing.
6. Run the acceptance tests `$ make testacc`
7. Create a PR and send it our way :)

#### Acceptance Test Environment Variables
Variable | Command | Description
---------|-------------|--------
TF_ACC|`$ export TF_ACC=1`|Enables online acceptance tests.
IC_USERNAME|`$ export IC_USERNAME=<your instaclustr username>`|Authorizes Provisioning API
IC_API_KEY|`$ export IC_API_KEY=<your provisioning API key>`|Authorizes Provisioning API
KMS_ARN|`$ export KMS_ARN=<your KMS ARN>`|For EBS encryption of nodes. <b><i>Note:</i></b> You cannot use an ARN previously added to your account as an encryption key.
IC_PROV_ACC_NAME|`$ export IC_PROV_ACC_NAME="<your provider name>"`|Your "Run In Your Own Account" account name.
IC_PROV_VPC_ID|`$ export IC_PROV_VPC_ID="<your AWS VPC ID>"`|For provisioning into a custom VPC.

#### Environment Variables Specific to Kafka Connect Acceptance Test

These environment variables are optional and only required when we want to do acceptance tests for Kafka Connect.
It is toggled by setting IC_TEST_KAFKA_CONNECT environment variable.

Variable | Command | Description
---------|-------------|--------
IC_TEST_KAFKA_CONNECT|`$ export IC_PROV_VPC_ID=1`|Enables acceptance tests for Kafka Connect.
IC_TARGET_KAFKA_CLUSTER_ID|`$ export IC_PROV_VPC_ID="<target kafka cluster ID>"`|For Kafka Connect connection information. See bundle options.
IC_AWS_ACCESS_KEY|`$ export IC_PROV_VPC_ID="<access key for the AWS S3 bucket>"`|For Kafka Connect connection information. See bundle options.
IC_AWS_SECRET_KEY|`$ export IC_PROV_VPC_ID="<secret key for the AWS S3 bucket>"`|For Kafka Connect connection information. See bundle options.
IC_S3_BUCKET_NAME|`$ export IC_PROV_VPC_ID="<AWS S3 bucket name>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_ACCOUNT_NAME|`$ export IC_PROV_VPC_ID="<account name for the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_ACCOUNT_KEY|`$ export IC_PROV_VPC_ID="<account key for the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_CONTAINER_NAME|`$ export IC_PROV_VPC_ID="<the name of the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_ENABLED_PROTOCOLS|`$ export IC_PROV_VPC_ID="<SSL enabled protocols>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_TRUSTSTORE_PASSWORD|`$ export IC_PROV_VPC_ID="<SSL truststore password>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_PROTOCOL|`$ export IC_PROV_VPC_ID="<SSL protocol>"`|For Kafka Connect connection information. See bundle options.
IC_SECURITY_PROTOCOL|`$ export IC_PROV_VPC_ID="<Security protocol>"`|For Kafka Connect connection information. See bundle options.
IC_SASL_MECHANISM|`$ export IC_PROV_VPC_ID="<SASL mechanism>"`|For Kafka Connect connection information. See bundle options.
IC_SASL_JAAS_CONFIG|`$ export IC_PROV_VPC_ID="<SASL JAAS config>"`|For Kafka Connect connection information. See bundle options.
IC_BOOTSTRAP_SERVERS|`$ export IC_PROV_VPC_ID="<bootstrap servers>"`|For Kafka Connect connection information. See bundle options.
IC_TRUSTSTORE|`$ export IC_PROV_VPC_ID="<Base64 encoding of the truststore jks>"`|For Kafka Connect connection information. See bundle options.

## Further information and documentation

This provider makes use of the Instaclustr API.  For further information including latest updates and value definitions, please see [the provisioning API documentation](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/).

Please see https://www.instaclustr.com/support/documentation/announcements/instaclustr-open-source-project-status/ for Instaclustr support status of this project.

# License

Apache2 - See the included LICENSE file for more details.
